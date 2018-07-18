package main

import (
	"blogo-server/server/infrastructure"
	"os"
)

func main() {
	infrastructure.Router.Run(os.Getenv("API_SERVER_PORT"))
}
