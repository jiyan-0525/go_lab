package main

import (
	"os"
	"fmt"
	"gee"
	"net/http"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [port]")
		return
	}
	port := os.Args[1]

	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	r.POST("/login", func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		fmt.Fprintf(w, "username = %q\n", req.Form.Get("username"))
		fmt.Fprintf(w, "password = %q\n", req.Form.Get("password"))
	})
	r.Run(":" + port)
}
