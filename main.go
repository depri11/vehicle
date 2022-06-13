package main

import (
	"log"
	"os"

	"github.com/depri11/vehicle/src/configs/command"
)

func main() {
	if err := command.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	// mainRoute, err := routers.SetupRouter()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var addrs string = "0.0.0.0:3000"

	// if pr := os.Getenv("PORT"); pr != "" {
	// 	addrs = "0.0.0.0:" + pr
	// }

	// log.Println("App running on " + addrs)

	// if err := http.ListenAndServe(addrs, mainRoute); err != nil {
	// 	return
	// }
}
