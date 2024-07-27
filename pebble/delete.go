package pebble

import (
	"fmt"
	"github.com/cockroachdb/pebble"
)

func delRange(opts Config) bool {
	db, err := pebble.Open(opts.Path, &pebble.Options{})
	if err != nil {
		return false
	}
	defer db.Close()
	startTime := fmt.Sprintf("%s_%d", opts.Key, opts.DTSJ)
	endTime := fmt.Sprintf("%s_%d", opts.Key, opts.GQSJ)
	if err = db.DeleteRange([]byte(startTime), []byte(endTime), &pebble.WriteOptions{}); err != nil {
		return false
	}
	return true
}

func delKey(opts Config) bool {
	db, err := pebble.Open(opts.Path, &pebble.Options{})
	if err != nil {
		return false
	}
	defer db.Close()
	if err = db.Delete([]byte(opts.Key), &pebble.WriteOptions{}); err != nil {
		return false
	}
	return true
}
