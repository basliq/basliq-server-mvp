package main

import (
	"fmt"
	httpserver "github.com/basliq/basliq-server-mvp/server/http"
)

// TODO - setup docker for postgres service
// TODO - dockerize the project itself
// TODO - add postgres driver
// TODO - add jwt
// TODO - add rich error
// TODO - add validator
// TODO - add config
// TODO - add migrator

func main() {
	httpserver.Serve()
	fmt.Println("Hello from backend guys!")
}
