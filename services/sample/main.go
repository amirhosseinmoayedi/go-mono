package main

import (
	"github.com/amirhosseinmoayedi/go-mono/services/sample/server"
	"net/http"
)

func main() {
	http.HandleFunc("/", server.HelloWorld)
	http.ListenAndServe(":8080", nil)
}
