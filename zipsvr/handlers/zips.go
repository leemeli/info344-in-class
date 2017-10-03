package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/leemeli/info344-in-class/zipsvr/models"
)

type CityHandler struct {
	PathPrefix string
	Index      models.ZipIndex
}

// Receiver functions = how GO kind of does object oriented functions
// Left side of this function is a pointer so that we can access zip index (this pointer in Java,
// except more explicit!)
func (ch *CityHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// The URL we want our server to support:
	// /zips/city-name
	cityName := r.URL.Path[len(ch.PathPrefix):] // this gives us the city-name part of the url
	cityName = strings.ToLower(cityName)        // makes it case insensitive
	if len(cityName) == 0 {
		// We do not want to do logFatal here because the server will crash
		http.Error(w, "please provide a city name", http.StatusBadRequest)
		// StatusBadRequest is status code 400
		return
	}

	w.Header().Add(headerContentType, contentTypeJSON)
	w.Header().Add(accessControlAllowOrigin, "*")
	zips := ch.Index[cityName] // We know this will give us a list of city zips for the city name
	json.NewEncoder(w).Encode(zips)
}
