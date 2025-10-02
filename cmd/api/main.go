package main //Programs start running in package main.

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")

}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "GET Request")

	case http.MethodPost:
		fmt.Fprintln(w, "POST Request")

	default:
		fmt.Fprintf(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/user", handleRequest)
	http.ListenAndServe(":8080", nil)
}
