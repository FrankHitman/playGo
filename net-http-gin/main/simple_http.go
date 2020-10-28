package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/info", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("info"))
	})

	http.ListenAndServe(":8888", nil)
}
