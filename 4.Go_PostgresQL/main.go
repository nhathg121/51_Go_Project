package main

import (
	"fmt"

	"github.com/nhathg121/51_Go_Project/4.Go%26PostgresQL/handler"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Hello, World!")
	e := echo.New()
	e.GET("/", handler.Welcome)

	// e.GET("/hello/:name", handler.Hello)
}
