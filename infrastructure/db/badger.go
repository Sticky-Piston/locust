package db

import (
	"log"

	badger "github.com/dgraph-io/badger/v3"
)

type BadgerImpl struct {
	DB *badger.DB
}

func NewBadgerClient() *BadgerImpl {
	log.Println("Initializing badger")

	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	db, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return &BadgerImpl{DB: db}
}
