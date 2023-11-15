package badger

import (
	"github.com/dgraph-io/badger/v4"
	"log"
)

func DB() {
	db, err := badger.Open(badger.DefaultOptions("d:/configuration/badger"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
