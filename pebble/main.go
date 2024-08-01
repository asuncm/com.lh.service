package pebble

import (
	"bytes"
	"encoding/gob"
	"errors"
	"time"
)

type DConf = map[string]interface{}

type Config struct {
	Path     string        `json:"path"`
	Key      string        `json:"key"`
	Max      int           `json:"max"`
	DTSJ     string        `json:"dtsj"`
	GQSJ     string        `json:"gqsj"`
	Duration time.Duration `json:"duration"`
	Data     DConf         `json:"data"`
}

func EnCode(data DConf) (bytes.Buffer, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(data); err != nil {
		return buf, errors.New("serialization")
	}
	return buf, nil
}

func DeCode(data string) (DConf, error) {
	var buf bytes.Buffer
	var lists DConf
	_, err := buf.WriteString(data)
	if err != nil {
		return lists, errors.New("serialization")
	}
	enc := gob.NewDecoder(&buf)

	if err := enc.Decode(&lists); err != nil {
		return lists, errors.New("serialization")
	}
	return lists, nil
}
