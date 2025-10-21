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
	e.Static("/app/static", "static")
	e.File("/", "static/index.html")
	e.File("/profile", "static/profile.html")
	e.File("/tasks", "static/tasks.html")
	e.File("/news", "static/news.html")
	e.File("/support", "static/support.html")
	e.File("/login", "static/login.html")

	e.GET("/h", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	return e, nil
}

func jwtMiddleware(c config.Config) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(c.Secret),
	})
}
