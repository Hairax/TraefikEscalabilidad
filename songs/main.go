package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, []map[string]string{
			{"title": "Bad Guy", "artist": "Billie Eilish"},
			{"title": "Old Town Road", "artist": "Lil Nas X"},
		})
	})
	e.Logger.Fatal(e.Start(":8080"))
}
