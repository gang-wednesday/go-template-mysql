package throttle

// import (
// 	"fmt"
// 	"go-template/pkg/utl/redis"

// 	"github.com/labstack/echo/v4"
// 	"golang.org/x/net/ipv4"
// )

// func Throttle(next echo.HandlerFunc)echo.HandlerFunc{
//     return func(c echo.Context) error {
//         path :=c.Request().URL.Path
//         ip := echo.ExtractIPDirect()
//         ipAddr :=fmt.Sprintf("ip: %v", ip)
//         key := ipAddr+"|"+path
//         redis.
//     }
// }
