package pebble

import (
	"bytes"
	"com.lh.service/tools"
	"encoding/gob"
	"errors"
	"github.com/cockroachdb/pebble"
	"log"
)

func Append(opts Config, res interface{}) error {
	db, err := OpenDB(opts.Path, opts.Config)
	var msg string
	if err != nil {
		msg = err.(string)
	} else {
		listData := make(Conf)
		if !tools.IsStringEmtpy(opts.Key) {
			val, closer, err := db.Get([]byte(opts.Key))
			if err != nil && !tools.IsStringEmtpy(opts.StartTime) {
				listData[opts.Prefix+"_startTime"] = opts.StartTime
			} else {
				list := gob.NewDecoder(bytes.NewReader(val))
				if err = list.Decode(&listData); err != nil {
					msg = "serialization"
				} else {
					opts.StartTime = listData[opts.Prefix+"_startTime"].(string)
					startTime := delRange(db, opts)
					listData[opts.Prefix+"_startTime"] = startTime
				}
				if err = closer.Close(); err != nil {
					log.Fatal(err)
				}
				buf, err := EnCode(listData)
				if err != nil {
					msg = err.Error()
				} else {
					if err = db.Set([]byte(opts.Key), buf.Bytes(), pebble.Sync); err != nil {
						msg = "serialization"
					}
				}
			}
		}

		resStr, err := EnCode(res)

		if err != nil {
			msg = err.Error()
		} else {
			if err = db.Set([]byte(opts.Prefix+"_"+opts.ID), resStr.Bytes(), pebble.Sync); err != nil {
				msg = "serialization"
			}
		}
	}
	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
	return errors.New(msg)
}
