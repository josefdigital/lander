package routes

import (
	"github.com/joegasewicz/gomek"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, d *gomek.Data) {
	templateData := make(gomek.Data)
	templateData["Title"] = "Welcome to JOSEF Digital - Single Page Site"
	templateData["SiteName"] = "JOSEF Digital"
	*d = templateData
}
