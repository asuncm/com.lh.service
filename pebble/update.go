package pebble

import (
	"com.lh.service/tools"
	"github.com/cockroachdb/pebble"
	"log"
	"time"
)

type Option struct {
	Key       string `json:"cache"`
	ID        string `json:"key"`
	Min       int16  `json:"min"`
	Prefix    string `json:"prefix"`
	StartTime string `json:"startTime"`
}

type Conf = map[string]interface{}

func mapMerge(list Conf, data Conf) Conf {
	for key, value := range data {
		list[key] = value
	}
	return list
}

type FN func(interface{}) (interface{}, error)

func Put(opts Config, data interface{}, fn FN) error {
	db, err := OpenDB(opts.Path, opts.Config)
	var msg interface{} = nil
	if err != nil {
		msg = err
	} else {
		val, closer, err := db.Get([]byte(opts.Prefix + "_" + opts.ID))
		if err != nil {
			msg = "get"
		} else {
			list, err := fn(val)
			if err != nil {
				msg = err.Error()
			} else {
				vals := list.(Conf)
				if !tools.IsInt64Emtpy(vals["ExpirationTime"].(int64)) {
					expirationTime := time.Unix(vals["ExpirationTime"].(int64), 0)
					nowTime := time.Now()
					if nowTime.Before(expirationTime) {
						data = mapMerge(vals, data.(Conf))
					} else {
						msg = "timeout"
						if err = delKey(db, opts.Prefix+"_"+opts.ID); err != nil {
							msg = err
						}
					}
				} else {
					data = mapMerge(vals, data.(Conf))
				}
			}
			if err = closer.Close(); err != nil {
				log.Fatal(err)
			}
			buf, err := EnCode(data)
			if err != nil {
				msg = err.Error()
			} else if err = db.Set([]byte(opts.Prefix+"_"+opts.ID), buf.Bytes(), pebble.Sync); err != nil {
				msg = "save"
			}
		}
	}
	return msg.(error)
}
