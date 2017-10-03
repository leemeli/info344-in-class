package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/leemeli/info344-in-class/zipsvr/handlers"
	"github.com/leemeli/info344-in-class/zipsvr/models"
)

const zipsPath = "/zips/"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, "Hello %s!", name)
}

func memoryHandler(w http.ResponseWriter, r *http.Request) {
	runtime.GC() // Garbage collector does cause a bit of a scalability issue
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(stats)
}

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":80"
	}
	zips, err := models.LoadZips("zips.csv")
	if err != nil {
		// do not do log.fatal in HTTP handlers
		log.Fatalf("error loading zips: %v", err)
	}
	log.Printf("loaded %d zips", len(zips))

	// Want to be able to return all the zip codes of Seattle (given that we don't know the order)
	// Best way is by making a map
	cityIndex := models.ZipIndex{}
	// Designing the underscore into the Go language makes the language more explicit
	// Won't have random variables that are never used
	for _, z := range zips {
		cityLower := strings.ToLower(z.City)
		cityIndex[cityLower] = append(cityIndex[cityLower], z)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler) // HandleFunc takes a simple function
	mux.HandleFunc("/memory", memoryHandler)

	cityHandler := &handlers.CityHandler{
		Index:      cityIndex,
		PathPrefix: zipsPath, // Need a comma after every line, unlike javascript
	}

	mux.Handle(zipsPath, cityHandler) // Handle handles the HTTP Handler interface

	fmt.Printf("server is listening at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
