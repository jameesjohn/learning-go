package main

import (
	"fmt"
	"net/http"
)

func middleware(originalHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request phase!")
		// Pass control back to the handler.
		originalHandler.ServeHTTP(w, r)
		fmt.Println("Executing middleware after response phase!")
	})
}

func handle(w http.ResponseWriter, r *http.Request) {
	// Business logic goes here
	fmt.Println("Executing mainHandler...")
	w.Write([]byte("OK"))
}

func main() {
	// HandlerFunc returns a HTTP Handler
	originalHandler := http.HandlerFunc(handle)
	http.Handle("/", middleware(originalHandler))
	http.ListenAndServe(":8000", nil)
	fmt.Println("Listening on port 8000")
}
