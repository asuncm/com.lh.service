package pebble

import (
	"bytes"
	"encoding/gob"
	"errors"
	"github.com/cockroachdb/pebble"
	"time"
)

type DConf = map[string]interface{}

type Config struct {
	Path     string         `json:"path"`
	Config   pebble.Options `json:"config"`
	Key      string         `json:"key"`
	Max      int            `json:"max"`
	DTSJ     int64          `json:"dtsj"`
	GQSJ     int64          `json:"gqsj"`
	Duration time.Duration  `json:"duration"`
	Data     DConf          `json:"data"`
}

func OpenDB(opts Config) (*pebble.DB, error) {
	//platform := tools.Platform("")
	db, err := pebble.Open(opts.Path, &opts.Config)
	if err != nil {
		return nil, errors.New("SQLite")
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
