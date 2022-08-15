package main

import (
	"fmt"
	"net/http"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
)

func main() {
	http.HandleFunc("/", HelloServer)
	fmt.Println("server listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	say, _ := cowsay.Say(
		"Hello",
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
	)
	fmt.Fprintf(w, say)
}
