package main

import (
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Hello, World!")
	e := echo.New()
	e.GET("/", handler.Welcome)

	// e.GET("/hello/:name", handler.Hello)
}
