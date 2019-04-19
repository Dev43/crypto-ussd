package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Post("/ussd", func(w http.ResponseWriter, r *http.Request) {
		sessionId := r.FormValue("sessionId")
		phoneNumber := r.FormValue("phoneNumber")
		networkCode := r.FormValue("networkCode")
		text := r.FormValue("text")

		fmt.Println(sessionId, phoneNumber, networkCode, text)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}
	fmt.Println("Listnening on port", port)
	http.ListenAndServe(":3000", r)
}
