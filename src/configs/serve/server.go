package serve

import (
	"log"
	"net/http"
	"os"

	"github.com/depri11/vehicle/src/routers"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "start api server",
	RunE:  Serve,
}
var validate *validator.Validate

func Serve(cmd *cobra.Command, args []string) error {
	if mainRoute, err := routers.SetupRouter(); err == nil {
		var addrs string = "0.0.0.0:3000"

		if pr := os.Getenv("PORT"); pr != "" {
			addrs = "0.0.0.0:" + pr
		}

		validate = validator.New()

		log.Println("App running on " + addrs)

		if err := http.ListenAndServe(addrs, mainRoute); err != nil {
			return err
		}

		return nil
	} else {
		return err
	}
}
