// demo-1

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Printf("Greeting from %s\n", os.Getenv("GREETING"))

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":80", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Greeting from %s", os.Getenv("GREETING"))
}