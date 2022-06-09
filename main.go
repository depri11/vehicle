package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/depri11/vehicle/src/routers"
)

func main() {
	mux, err := routers.SetupRouter()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	fmt.Println("Running on port", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal("Failed running server")
	}

}
