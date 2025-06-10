package server

import (
	"portal/pkg/config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitServer(c config.Config) (*echo.Echo, error) {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
			echo.HeaderXRealIP,
			echo.HeaderXForwardedFor,
			echo.HeaderXRequestedWith,
		},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
	}))

	return e, nil
}

func jwtMiddleware(c config.Config) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(c.Secret),
	})
}
