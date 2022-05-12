package main

import (
	"net/http"
	"simple-web-framework/framework"
)

func main() {
	r := framework.New()

	r.GET("/index", func(c *framework.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *framework.Context) {
			c.HTML(http.StatusOK, "<h1>Hello world</h1>")
		})

		v1.GET("/hello", func(c *framework.Context) {
			// expect /hello?name=cooper
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *framework.Context) {
			// expect /hello/cooper
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *framework.Context) {
			c.JSON(http.StatusOK, framework.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	err := r.Run("localhost:3096")
	if err != nil {
		return
	}
}
