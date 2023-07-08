package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TechnoGeek01/learning-go/config"
	"github.com/TechnoGeek01/learning-go/pkg/handlers"
	"github.com/TechnoGeek01/learning-go/pkg/render"
)

const portNumber = ":8080"

// Main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)

}
