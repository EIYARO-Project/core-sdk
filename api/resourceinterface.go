package api

import "github.com/EIYARO-Project/core-sdk/api/resources"

type Resources interface {
	resources.AccessToken |
		resources.Account |
		any
}

type ResourceInterface[T Resources] interface {
	List() ([]T, error)
	//View() (T, error)
	// Edit() (T, error)
	// Add() (T, error)
	// Delete() (bool, error)
}
