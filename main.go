package main

import (
	"github.com/joegasewicz/gomek"
	"github.com/josefdigital/lander/routes"
	"net/http"
)

func main() {
	c := gomek.Config{
		BaseTemplateName: "layout",
		BaseTemplates: []string{
			"templates/layout.gohtml",
			"templates/partials/navbar.gohtml",
			"templates/partials/footer.gohtml",
			"templates/partials/hero-section.gohtml",
			"templates/partials/info-section.gohtml",
			"templates/partials/product-section.gohtml",
			"templates/partials/what-we-do-section.gohtml",
			"templates/partials/service-section.gohtml",
			"templates/partials/connect-section.gohtml",
			"templates/partials/contact-us-section.gohtml",
			"templates/partials/product-card.gohtml",
			"templates/partials/what-we-do-card.gohtml",
			"templates/partials/service-card.gohtml",
		},
	}
	app := gomek.New(c)

	// Static files
	staticFiles := http.FileServer(http.Dir("static"))
	app.Handle("/static/", http.StripPrefix("/static/", staticFiles))

	app.
		Route("/").
		View(routes.HomeHandler).
		Methods("GET", "POST").
		Templates("templates/routes/index.gohtml")

	app.Use(gomek.Logging)
	app.Use(gomek.CORS)

	app.Listen(6011)
	app.Start()
}
