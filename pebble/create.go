package pebble

import (
	"fmt"
	"github.com/cockroachdb/pebble"
	"strconv"
	"time"
)

func Add(opts Config) (DConf, error) {
	db, err := pebble.Open(opts.Path, &pebble.Options{})
	if err != nil {
		return DConf{}, err
	}
	defer db.Close()
	id := fmt.Sprintf("%s_%d", opts.Key, opts.DTSJ)
	opts.Data["id"] = id
	buf, err := EnCode(opts.Data)
	if err != nil {
		return DConf{}, err
	}
	val := []byte(buf.String())
	if err := db.Set([]byte(id), val, pebble.Sync); err != nil {
		return DConf{}, err
	}
	return opts.Data, nil
}

func AddLog(opts Config) (DConf, error) {
	db, err := pebble.Open(opts.Path, &pebble.Options{})
	if err != nil {
		return DConf{}, err
	}
	defer db.Close()
	start, closer, err := db.Get([]byte(fmt.Sprintf("%s_%s", opts.Key, "time")))
	var startTime int64
	nowTime := time.Now()
	endTime := nowTime.Add(-opts.Duration).UnixNano()
	if err != nil {
		startTime = endTime
	} else {
		startT, err := strconv.ParseInt(string(start), 10, 64)
		if err != nil {
			startTime = endTime
		} else {
			startTime = startT
		}
		defer closer.Close()
	}
	_ = db.DeleteRange([]byte(fmt.Sprintf("%s_%d", opts.Key, startTime)), []byte(fmt.Sprintf("%s_%d", opts.Key, endTime)), &pebble.WriteOptions{})
	id := fmt.Sprintf("%s_%d", opts.Key, opts.DTSJ)
	opts.Data["id"] = id
	buf, err := EnCode(opts.Data)
	if err != nil {
		return DConf{}, err
	}
	val := []byte(buf.String())
	if err := db.Set([]byte(id), val, pebble.Sync); err != nil {
		return DConf{}, err
	}
	_ = db.Set([]byte(fmt.Sprintf("%s_%s", opts.Key, "time")), []byte(string(endTime)), pebble.Sync)
	return opts.Data, nil
}
