package pebble

import (
	"bytes"
	"encoding/gob"
	"errors"
	"github.com/cockroachdb/pebble"
)

type Config struct {
	Path   string         `json:"path"`
	Config pebble.Options `json:"config"`
	Option
}

func OpenDB(pathname string, opts pebble.Options) (*pebble.DB, interface{}) {
	db, err := pebble.Open(pathname, &opts)
	if err != nil {
		str := "SQLite"
		db.Close()
		return nil, str
	} else {
		return db, nil
	}
}

func EnCode(data interface{}) (bytes.Buffer, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(data); err != nil {
		return buf, errors.New("serialization")
	} else {
		return buf, nil
	}
}
