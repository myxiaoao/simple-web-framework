package main

import (
	"net/http"
	"simple-web-framework/framework"
)

func main() {
	w := framework.New()

	w.GET("/", func(c *framework.Context) {
		c.HTML(http.StatusOK, "<h1>Hello world.</h1>")
	})
	w.GET("/hello", func(c *framework.Context) {
		// expect /hello?name=cooper
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	w.POST("/login", func(c *framework.Context) {
		c.JSON(http.StatusOK, framework.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	err := w.Run("localhost:3096")
	if err != nil {
		return
	}
}
