// Package api Go-Template
package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	graphql "go-template/gqlmodels"
	"go-template/internal/config"
	"go-template/internal/jwt"
	authMw "go-template/internal/middleware/auth"
	"go-template/internal/middleware/throttle"
	"go-template/internal/mysql"
	"go-template/internal/server"
	"go-template/resolver"

	"go-template/internal/secondary"

	graphql2 "github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	echo "github.com/labstack/echo/v4"
	_ "github.com/lib/pq" // here
	"github.com/sony/gobreaker"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Start starts the API service
func Start(cfg *config.Configuration) (*echo.Echo, error) {
	fmt.Printf("connection started")

	db, err := mysql.Connect()
	if err != nil {
		return nil, err
	}

	boil.SetDB(db)

	jwt, err := jwt.New(
		cfg.JWT.SigningAlgorithm,
		os.Getenv("JWT_SECRET"),
		cfg.JWT.DurationMinutes,
		cfg.JWT.MinSecretLength)
	if err != nil {
		return nil, err
	}

	e := server.New()

	gqlMiddleware := authMw.GqlMiddleware()
	// throttlerMiddleware puts the current user's IP address into context of gqlgen
	// throttlerMiddleware :

	graphQLPathname := "/graphql"
	playgroundHandler := playground.Handler("GraphQL playground", graphQLPathname)

	graphqlHandler := handler.New(graphql.NewExecutableSchema(graphql.Config{
		Resolvers: &resolver.Resolver{},
	}))

	if os.Getenv("ENVIRONMENT_NAME") == "local" {
		boil.DebugMode = true
	}

	// graphql apis
	graphqlHandler.AroundOperations(func(ctx context.Context, next graphql2.OperationHandler) graphql2.ResponseHandler {
		return authMw.GraphQLMiddleware(ctx, jwt, next)
	})

	e.POST(graphQLPathname, func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		graphqlHandler.ServeHTTP(res, req)
		return nil
	}, gqlMiddleware, throttle.ThrottlingMiddleware)

	e.GET(graphQLPathname, func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		graphqlHandler.ServeHTTP(res, req)
		return nil
	}, gqlMiddleware, throttle.ThrottlingMiddleware)

	graphqlHandler.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	graphqlHandler.AddTransport(transport.Options{})
	graphqlHandler.AddTransport(transport.GET{})
	graphqlHandler.AddTransport(transport.POST{})
	graphqlHandler.AddTransport(transport.MultipartForm{})

	graphqlHandler.SetQueryCache(lru.New(1000))

	graphqlHandler.Use(extension.Introspection{})
	graphqlHandler.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	var st gobreaker.Settings
	st.Name = "HTTP GET"
	st.ReadyToTrip = func(counts gobreaker.Counts) bool {
		failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
		return counts.Requests >= 3 && failureRatio >= 0.6
	}

	secondary.Cb = gobreaker.NewCircuitBreaker(st)

	// graphql playground
	e.GET("/playground", func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		playgroundHandler.ServeHTTP(res, req)
		return nil
	})
	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})
	return e, nil
}
