package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"

	d "github.com/fw8/google-auth-proxy/app/director"
	rt "github.com/fw8/google-auth-proxy/app/router"
)

var rp = httputil.ReverseProxy{
	Director: d.NewDirector(),
}

func main() {

	fmt.Printf("SCHEME %v\n", os.Getenv("SCHEME"))
	fmt.Printf("HOST %v\n", os.Getenv("HOST"))
	fmt.Printf("GOOGLE_APPLICATION_CREDENTIALS %v\n", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

	fmt.Println("Starting proxy server")

	r := rt.NewRouter(&rp)

	srv := http.Server{
		Handler: r,
		Addr:    ":8080",
	}
	srv.ListenAndServe()
}
