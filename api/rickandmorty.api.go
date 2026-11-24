package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"guide/models"
)

// URL de base de l'API et identification de notre programme.
var (
	baseURL   = "themealdb.com/api/json/v1/1/categories.php"
	userAgent = "RickAndMorty-Ynov/1.0 (contact@exemple.fr)"
)

// Un seul client HTTP, partagé par toutes les requêtes.
var client = &http.Client{
	Timeout: 5 * time.Second,
}

// get envoie une requête GET vers l'endpoint et décode la réponse dans dest.
func get(endpoint string, dest any) error {
	// 1. Préparer la requête : méthode, URL et corps (nil pour un GET).
	request, err := http.NewRequest(http.MethodGet, baseURL+endpoint, nil)
	if err != nil {
		return fmt.Errorf("création de la requête : %w", err)
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", userAgent)

	// 2. Envoyer la requête avec le client.
	response, err := client.Do(request)
	if err != nil {
		return fmt.Errorf("%w : %v", models.ErrUnavailable, err)
	}
	defer response.Body.Close()

	// 3. Vérifier le code de statut avant de lire le corps.
	switch {
	case response.StatusCode == http.StatusOK:
		// 200 : succès, on peut décoder le corps.
	case response.StatusCode == http.StatusNotFound:
		return models.ErrNotFound
	case response.StatusCode == http.StatusBadRequest:
		return models.ErrBadRequest
	case response.StatusCode >= 500:
		return fmt.Errorf("%w : %s", models.ErrUnavailable, response.Status)
	default:
		return fmt.Errorf("réponse inattendue : %s", response.Status)
	}

	// 4. Décoder le flux JSON directement dans la variable de destination.
	if err := json.NewDecoder(response.Body).Decode(dest); err != nil {
		return fmt.Errorf("décodage de la réponse : %w", err)
	}
	return nil
}

// GetCharacters renvoie la première page des personnages (endpoint /character).
// Si name n'est pas vide, l'API ne renvoie que les personnages dont le nom le contient.
func GetCharacters(name string) (models.CharacterPage, error) {
	endpoint := "/character"
	if name != "" {
		params := url.Values{}
		params.Set("name", name)
		endpoint += "?" + params.Encode()
	}
	var page models.CharacterPage
	err := get(endpoint, &page)
	return page, err
}

// GetCharacter renvoie la fiche d'un personnage (endpoint /character/{id}).
func GetCharacter(id int) (models.Character, error) {
	var character models.Character
	err := get(fmt.Sprintf("/character/%d", id), &character)
	return character, err
}
