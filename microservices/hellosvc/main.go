package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type HelloHandler struct {
	serviceAddr string
}

func (th *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello! The time service at %s says the current time is %s",
		th.serviceAddr, time.Now().Format(time.Kitchen))
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}

	http.Handle("/v1/hello", &HelloHandler{addr})
	log.Printf("server is listening at http://%s...", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
