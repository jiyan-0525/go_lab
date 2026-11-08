package main

import (
	"fmt"
	"log"
	"net/http"
	"os"     // 必须导入这个包来使用 Getenv
	"time"
)

func main() {
	// 1. 获取环境变量 PORT
	port := os.Getenv("PORT")

	// 2. 如果环境变量没有设置，可以给一个默认值（可选，但作业要求设为 8999）
	if port == "" {
		port = "8010" 
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Docker! I'm a Go server.")
	})

	server := &http.Server{
		Addr:         ":" + port, // 这里会变成 ":8999"
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("server started on %s\n", port)
	log.Fatal(server.ListenAndServe())
}