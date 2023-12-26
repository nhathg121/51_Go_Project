package user_handler

import "github.com/labstack/echo"

func Welcome(c echo.Context) error {
	return c.String(200, "Welcome to my website!")
}
