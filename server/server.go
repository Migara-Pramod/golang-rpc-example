package main

import (
	"io"
	"net/http"
	"net/rpc"
	"rpc.example.com/servercore"
)

func main() {

	// create a `*Vegetables` object
	vegetables := servercore.NewVegetables()
	
	// register `vegetables` object with `rpc.DefaultServer`
	rpc.Register(vegetables)

	// register an HTTP handler for RPC communication on `http.DefaultServeMux` (default)
	// registers a handler on the `rpc.DefaultRPCPath` endpoint to respond to RPC messages
	// registers a handler on the `rpc.DefaultDebugPath` endpoint for debugging
	rpc.HandleHTTP()

	// sample test endpoint
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "RPC SERVER LIVE!")
	})

	// listen and serve default HTTP server
	http.ListenAndServe(":9000", nil)

}
