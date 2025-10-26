package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Static("/static", "assets")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/save", save)
	e.POST("/users", func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
		// or
		// return c.XML(http.StatusCreated, u)
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	fmt.Println(member)
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}
func save(c echo.Context) error {
	// Get name and email
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.String(http.StatusOK, "name:"+u.Name+", email:"+u.Email)
}
