package main

import (
	"net/http"
	"simple-web-framework/framework"
)

func main() {
	w := framework.New()

	w.GET("/", func(c *framework.Context) {
		c.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	w.GET("/hello", func(c *framework.Context) {
		// expect /hello?name=cooper
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	w.GET("/hello/:name", func(c *framework.Context) {
		// expect /hello/cooper
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	w.GET("/assets/*filepath", func(c *framework.Context) {
		c.JSON(http.StatusOK, framework.H{"filepath": c.Param("filepath")})
	})

	err := w.Run("localhost:3096")
	if err != nil {
		return
	}
}
