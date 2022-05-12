package main

import (
	"log"
	"net/http"
	"simple-web-framework/framework"
	"time"
)

func onlyForV2() framework.HandlerFunc {
	return func(c *framework.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := framework.New()

	r.Use(framework.Logger()) // global middleware
	r.GET("/", func(c *framework.Context) {
		c.HTML(http.StatusOK, "<h1>Hello world</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *framework.Context) {
			// expect /hello/cooper
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	err := r.Run("localhost:3096")
	if err != nil {
		return
	}
}
