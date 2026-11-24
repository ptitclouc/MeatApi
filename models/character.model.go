package models

// CharacterPage : une page de la liste des personnages.
// Seul le tableau "results" est utile ici (l'objet "info" sert à la pagination).
type CharacterPage struct {
	Results []Character `json:"results"`
}

// Character : fiche d'un personnage renvoyée par The Rick and Morty API.
// Seuls les champs utiles au site sont déclarés.
type Character struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Status   string   `json:"status"`
	Species  string   `json:"species"`
	Gender   string   `json:"gender"`
	Image    string   `json:"image"`
	Origin   Place    `json:"origin"`
	Location Place    `json:"location"`
	Episode  []string `json:"episode"`
}

// Place : un lieu (origine ou lieu actuel du personnage).
type Place struct {
	Name string `json:"name"`
}
