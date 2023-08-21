package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Inside the / path")
	})

	fmt.Println("Litening on port 8080")
	http.ListenAndServe(":8080", nil)

}
