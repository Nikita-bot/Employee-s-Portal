package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitServer() (*echo.Echo, error) {
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
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	return e, nil
}
