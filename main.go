package main

import (
	"fmt"
	"net/http"
	"simple-web-framework/framework"
)

func main() {
	w := framework.New()

	w.GET("/", func(w http.ResponseWriter, req *http.Request) {
		_, err := fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
		if err != nil {
			return
		}
	})

	w.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			_, err := fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
			if err != nil {
				return
			}
		}
	})

	err := w.Run("localhost:3096")
	if err != nil {
		return
	}
}
