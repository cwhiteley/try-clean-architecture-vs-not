package main

import (
	"blogo-server/api/infrastructure"
	"os"
)

func main() {
	infrastructure.Router.Run(os.Getenv("API_SERVER_PORT"))
}
