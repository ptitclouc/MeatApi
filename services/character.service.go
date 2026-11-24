package services

import (
	"errors"
	"strconv"

	"guide/api"
	"guide/models"
)

// SearchCharacters renvoie les personnages de la première page,
// filtrés par l'API sur le nom si une recherche est faite.
func SearchCharacters(name string) ([]models.Character, error) {
	page, err := api.GetCharacters(name)
	// Une recherche sans résultat n'est pas une erreur pour notre site :
	// l'API répond 404, on renvoie simplement une liste vide.
	if errors.Is(err, models.ErrNotFound) && name != "" {
		return []models.Character{}, nil
	}
	if err != nil {
		return nil, err
	}
	return page.Results, nil
}

// GetCharacter vérifie l'identifiant reçu dans l'URL, puis appelle l'API.
func GetCharacter(idParam string) (models.Character, error) {
	id, err := strconv.Atoi(idParam)
	if err != nil || id < 1 {
		return models.Character{}, models.ErrBadRequest
	}
	return api.GetCharacter(id)
}
