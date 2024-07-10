package pebble

import (
	"com.lh.service/tools"
	"errors"
	"github.com/cockroachdb/pebble"
	"strconv"
	"time"
)

func delRange(db *pebble.DB, opts Config) string {
	if tools.IsInt16Emtpy(opts.Min) {
		start := opts.StartTime
		num, err := strconv.ParseInt(start, 10, 64)
		if err == nil {
			startTime := time.Unix(num, 0)
			now := time.Now()
			endTime := now.Add(time.Duration(0-opts.Min) * time.Minute)
			SN := opts.Prefix + "_" + string(startTime.UnixNano())
			EN := opts.Prefix + "_" + string(endTime.UnixNano())
			err := db.DeleteRange([]byte(SN), []byte(EN), nil)
			if err == nil {
				return endTime.Format("200601021504")
			}
		}
	}
	return opts.StartTime
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
