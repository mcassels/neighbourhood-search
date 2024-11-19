package main

import (
	"context"
	"fmt"
	"neighbourhood-search/internal/backend"
	"neighbourhood-search/internal/generate"
	"neighbourhood-search/internal/middleware"
	"neighbourhood-search/internal/template"
	"neighbourhood-search/internal/types"
	"neighbourhood-search/internal/view"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func CheckValidityHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CheckValidityHandler")
	fmt.Println(r.Form)
	template.SubmitButton(r.FormValue("text")).Render(context.Background(), w)
	// middleware.Chain(w, r, template.TestString(("new inner html")))
}

func GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetMessagesHandler")
	fmt.Println(r.FormValue("text"))
	template.SubmitButton(r.FormValue("text"))
}

func main() {

	err := generate.GenerateMain()
	if err != nil {
		panic(err)
	}

	_ = godotenv.Load()
	mux := http.NewServeMux()

	var addresses = make([]types.AddressResult, 0)
	googleMapsAPIKey := os.Getenv("GOOGLE_MAPS_KEY")
	if googleMapsAPIKey == "" {
		fmt.Println("GOOGLE_MAPS_KEY is not set")
		return
	}

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		middleware.Chain(w, r, template.Home("Neighbourhood Finder", addresses))
	})

	mux.HandleFunc("GET /favicon.ico", view.ServeFavicon)
	mux.HandleFunc("GET /static/", view.ServeStaticFiles)
	mux.HandleFunc("POST /check-input", func(w http.ResponseWriter, r *http.Request) {
		middleware.Chain(w, r, template.SubmitButton(r.FormValue("text")))
	})
	mux.HandleFunc("POST /submit", func(w http.ResponseWriter, r *http.Request) {
		address := r.FormValue("text")

		latlng := ""
		for _, a := range addresses {
			if a.Address == address {
				latlng = a.LatLng
				break
			}
		}
		if latlng == "" {
			latlng = backend.Geocode(googleMapsAPIKey, address)
		}
		fmt.Println(latlng)
		newAddress := types.AddressResult{
			Address:       r.FormValue("text"),
			LatLng:        latlng,
			Neighbourhood: "this is a test",
		}

		addresses = append(addresses, newAddress)
		middleware.Chain(w, r, template.AddressRow(newAddress))
	})

	port := os.Getenv("PORT")
	fmt.Printf("server is running on port %s\n", port)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		fmt.Println(err)
	}

}
