package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("GO Echo Server")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Yallo from the other side")
	})
	e.Start(":8080")
}
