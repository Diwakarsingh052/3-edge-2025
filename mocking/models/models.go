package models

// run go generate command from the cli
// make sure terminal is in the current location

// mockgen would create mock implementation of the interface

//go:generate mockgen -source models.go -destination mockmodels/models_mock.go -package mockmodels
type DB interface {
	create(s string)
	update(id int)
	delete(id int)
}
