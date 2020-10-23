package main

import (
	"github.com/ashishkhuraishy/blogge/src/app"
	"github.com/ashishkhuraishy/blogge/src/docs"
)

func main() {

	docs.SwaggerInfo.Title = "Blogge API"
	docs.SwaggerInfo.Description = "This is a sample server Blogge server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	app.StartApplication()
}
