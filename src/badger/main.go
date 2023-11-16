package badger

import (
	"github.com/dgraph-io/badger/v4"
	"log"
)

func DB(pathname string) {
	db, err := badger.Open(badger.DefaultOptions(pathname))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
