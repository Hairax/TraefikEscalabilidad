package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, []map[string]string{
			{"title": "The Catcher in the Rye", "author": "J.D. Salinger"},
			{"title": "To Kill a Mockingbird", "author": "Harper Lee"},
		})
	})
	e.Logger.Fatal(e.Start(":8080"))
}
