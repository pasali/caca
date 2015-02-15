// Package caca provides simple HTTP server
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var dir string
var port string

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Method: %s URL: %s\n", r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func init() {
	pwd, _ := os.Getwd()
	flag.StringVar(&dir, "dir", pwd, "path of target directory")
	flag.StringVar(&port, "port", "8080", "destination port number")
	flag.Parse()
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(dir)))
	http.ListenAndServe(":"+port, Log(http.DefaultServeMux))
}
