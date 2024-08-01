package pebble

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/cockroachdb/pebble"
	"time"
)

func Add(opts Config) (DConf, error) {
	db, err := pebble.Open(opts.Path, &pebble.Options{})
	if err != nil {
		return DConf{}, err
	}
	defer db.Close()
	prefix := fmt.Sprintf("%s_%s", opts.Key, opts.DTSJ)
	iter, err := db.NewIter(&pebble.IterOptions{
		LowerBound: []byte(prefix),
		UpperBound: append([]byte(prefix), 0xff),
	})
	if err != nil {
		return DConf{}, errors.New("query")
	}
	count := 0
	for iter.First(); iter.Valid(); iter.Next() {
		key := iter.Key()
		count++
		if !bytes.HasPrefix(key, []byte(prefix)) {
			break
		}
	}
	id := fmt.Sprintf("%s%d", prefix, count)
	opts.Data["id"] = id
	buf, err := EnCode(opts.Data)
	if err != nil {
		return DConf{}, err
	}
	val := []byte(buf.String())
	if err := db.Set([]byte(id), val, pebble.Sync); err != nil {
		return DConf{}, errors.New("save")
	}
	return opts.Data, nil
}

func AddLog(opts Config) (DConf, error) {
	db, err := pebble.Open(opts.Path, &pebble.Options{})
	if err != nil {
		return DConf{}, errors.New("SQLite")
	}

	start, closer, err := db.Get([]byte(fmt.Sprintf("%s_%s", opts.Key, "time")))
	var startTime string
	nowTime := time.Now()
	endTime := nowTime.Add(-opts.Duration).Format("20060102150405")
	if err != nil {
		startTime = endTime
	} else {
		startTime = string(start)
		defer closer.Close()
	}
	_ = db.DeleteRange([]byte(fmt.Sprintf("%s_%s%d", opts.Key, startTime, 0)), []byte(fmt.Sprintf("%s_%s%d", opts.Key, endTime, 0)), &pebble.WriteOptions{})
	prefix := fmt.Sprintf("%s_%s", opts.Key, opts.DTSJ)
	fix := []byte(prefix)
	iter, err := db.NewIter(&pebble.IterOptions{
		LowerBound: fix,
		UpperBound: append(fix, 0xff),
	})
	if err != nil {
		return DConf{}, errors.New("query")
	}
	count := 0
	for iter.First(); iter.Valid(); iter.Next() {
		key := iter.Key()
		count++
		if !bytes.HasPrefix(key, fix) {
			break
		}
	}
	id := fmt.Sprintf("%s%d", prefix, count)
	opts.Data["id"] = id
	buf, err := EnCode(opts.Data)
	if err != nil {
		return DConf{}, err
	}
	val := []byte(buf.String())
	if err := db.Set([]byte(id), val, pebble.Sync); err != nil {
		return DConf{}, errors.New("save")
	}
	if err = db.Set([]byte(fmt.Sprintf("%s_%s", opts.Key, "time")), []byte(string(endTime)), pebble.Sync); err != nil {
		return DConf{}, err
	}
	defer db.Close()
	return opts.Data, nil
}
