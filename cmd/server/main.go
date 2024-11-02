package main

import (
	"net/http"

	"github.com/cffmnk/yametrics/internal/server"
)

func main() {
	server := server.NewServer()
	mux := http.NewServeMux()
	mux.HandleFunc("/update/", server.HandleUpdateMetrics)
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		panic(err)
	}
}
