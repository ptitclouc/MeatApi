package controllers

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"guide/helpers"
	"guide/models"
	"guide/services"
)

// ListPage : données de la page de liste.
type ListPage struct {
	Search     string // recherche saisie, pour la réafficher dans le formulaire
	Characters []models.Character
}

// ListCharacters affiche la liste des personnages, filtrée par ?nom=… si besoin.
func ListCharacters(w http.ResponseWriter, r *http.Request) {
	search := strings.TrimSpace(r.URL.Query().Get("nom"))
	characters, err := services.SearchCharacters(search)
	if err != nil {
		renderAPIError(w, err)
		return
	}
	helpers.Render(w, http.StatusOK, "list", ListPage{Search: search, Characters: characters})
}

// ShowCharacter affiche la fiche du personnage dont l'identifiant est dans l'URL.
func ShowCharacter(w http.ResponseWriter, r *http.Request) {
	character, err := services.GetCharacter(r.PathValue("id"))
	if err != nil {
		renderAPIError(w, err)
		return
	}
	helpers.Render(w, http.StatusOK, "character", character)
}

// NotFound affiche la page 404 pour toute URL inconnue du site.
func NotFound(w http.ResponseWriter, r *http.Request) {
	helpers.RenderError(w, http.StatusNotFound, "Cette page n'existe pas.")
}

// renderAPIError traduit une erreur connue en page d'erreur, sur la route même.
func renderAPIError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, models.ErrNotFound):
		helpers.RenderError(w, http.StatusNotFound,
			"Ce personnage n'existe pas.")
	case errors.Is(err, models.ErrBadRequest):
		helpers.RenderError(w, http.StatusBadRequest,
			"Identifiant de personnage invalide.")
	case errors.Is(err, models.ErrUnavailable):
		log.Println("The Rick and Morty API :", err)
		helpers.RenderError(w, http.StatusServiceUnavailable,
			"Service momentanément indisponible.")
	default:
		log.Println("Erreur inattendue :", err)
		helpers.RenderError(w, http.StatusInternalServerError,
			"Une erreur est survenue.")
	}
}
