package models

import "errors"

// Erreurs connues, partagées par les couches api, services et controllers.
var (
	ErrNotFound    = errors.New("ressource introuvable")
	ErrBadRequest  = errors.New("demande invalide")
	ErrUnavailable = errors.New("API indisponible")
)
