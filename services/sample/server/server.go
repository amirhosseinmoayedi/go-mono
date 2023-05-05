package server

import (
	"fmt"
	"net/http"
)

func HelloWorld(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "hell wordl!")
}
