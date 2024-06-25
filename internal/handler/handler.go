package handler

import (
	"crypto/subtle"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sinameshkini/webshare/internal/utils"
	"net/http"
)

type Handler struct{}

var (
	h *Handler
)

func Init(e *echo.Echo, path string, port int, username, password string) {
	h = &Handler{}

	// basic auth
	if username != "" && password != "" {
		e.Use(middleware.BasicAuth(func(usr, pass string, c echo.Context) (bool, error) {
			// Be careful to use constant time comparison to prevent timing attacks
			if subtle.ConstantTimeCompare([]byte(usr), []byte(username)) == 1 &&
				subtle.ConstantTimeCompare([]byte(pass), []byte(password)) == 1 {
				return true, nil
			}
			return false, nil
		}))

		fmt.Printf("Access protected by Basic Auth (Username: %s\tPassword: %s)\n", username, password)
	}

	e.Static("/", path)
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   path,
		Browse: true,
	}))
	e.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "status: running")
	})

	fmt.Printf("Enter this address in other host in local network: http://%s:%d \n", utils.GetIP(), port)
}
