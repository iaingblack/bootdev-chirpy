package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
}

type Server struct {
	mux *http.ServeMux
}
