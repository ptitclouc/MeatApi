package routes

import (
	"net/http"

	"guide/controllers"
)

// characterRoutes enregistre les routes de la partie personnages.
func characterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /personnages", controllers.ListCharacters)
	mux.HandleFunc("GET /personnages/{id}", controllers.ShowCharacter)
}
