package main

import (
	"fmt"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("uri: %s\n", req.RequestURI)
	res.WriteHeader(200)
	_, _ = res.Write([]byte("Hello World"))
}

func main() {
	sv := new(server)
	_ = http.ListenAndServe("0.0.0.0:8080", sv)
}
