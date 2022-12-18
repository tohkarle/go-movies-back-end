// This declares that this file is part of the 'main' package(?).
package main

// These are the imports from the standard library.
import (
	"fmt"
	"log"
	"net/http"
)

// This application is going to listen on port 8080.
const port = 8080

type application struct {
	Domain string
}

func main() {
	// Set application config.
	var app application

	// Read from command line.

	// Connect to the database.

	app.Domain = "example.com"

	log.Println("Starting application on port", port)

	http.HandleFunc("/", app.Home)

	// Start a web server.
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}

}