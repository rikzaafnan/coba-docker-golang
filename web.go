package main

import (
	"fmt"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome To Web Kita.")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Selamat datang di Web Kita")
	})

	http.HandleFunc("/welcome", welcome)

	http.ListenAndServe(":8080", nil)

}
