package throttle

import (
	redisutil "go-template/pkg/utl/redisUtil"

	"github.com/labstack/echo/v4"
)

func ThrottlingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		redisClient := redisutil.GetClient()

		counter, err := redisutil.GetCounter(c.Request().Context(), redisClient, c.RealIP())
		if err != nil {

			return err
		}
		err = redisutil.SetCounter(c.Request().Context(), redisClient, c.RealIP(), counter+1)
		if err != nil {

			return err
		}

		err = next(c)
		return err
	}
}
