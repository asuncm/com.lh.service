package pebble

import (
	"errors"
	"github.com/cockroachdb/pebble"
)

func delRange(db *pebble.DB, opts Config) string {

	return ""
}

func delKey(db *pebble.DB, key string) error {
	err := db.Delete([]byte(key), nil)
	if err != nil {
		return errors.New("del")
	}
	return nil
}

func Delkey(db *pebble.DB, key string) error {
	err := db.Delete([]byte(key), nil)
	if err != nil {
		return errors.New("del")
	}
	return nil
}
