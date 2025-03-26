package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, []map[string]string{
			{"title": "The Godfather", "director": "Francis Ford Coppola"},
			{"title": "The Shawshank Redemption", "director": "Frank Darabont"},
			{"title": "The Dark Knight", "director": "Christopher Nolan"},
		})
	})
	e.Logger.Fatal(e.Start(":8080"))
}
