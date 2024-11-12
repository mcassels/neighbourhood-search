package main

import (
	"context"
	"fmt"
	"neighbourhood-search/internal/generate"
	"neighbourhood-search/internal/middleware"
	"neighbourhood-search/internal/template"
	"neighbourhood-search/internal/view"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func CheckValidityHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CheckValidityHandler")
	fmt.Println(r.FormValue("text"))
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

	mux.HandleFunc("GET /favicon.ico", view.ServeFavicon)
	mux.HandleFunc("GET /static/", view.ServeStaticFiles)
	mux.HandleFunc("POST /check-validity", CheckValidityHandler)
	mux.HandleFunc("GET /get-message", GetMessagesHandler)

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		middleware.Chain(w, r, template.Home("Neighbourhood Finder"))
	})

	fmt.Printf("server is running on port %s\n", os.Getenv("PORT"))
	err = http.ListenAndServe(":"+os.Getenv("PORT"), mux)
	if err != nil {
		fmt.Println(err)
	}

}
