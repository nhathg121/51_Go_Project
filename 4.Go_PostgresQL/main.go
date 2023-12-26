package main

import (
	"fmt"
	"user_handler"
	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Hello, World!")
	e := echo.New()
	e.GET("/", user_handler.Welcome)
	e.GET("/hello/:name", handler.Hello)
	e.Logger.Fatal(e.Start(":1323"))
}
