package badger

import (
	bg "github.com/dgraph-io/badger/v4"
	"time"
)

func Append(pathname string, key string, data DataConf) error {
	db, err := Open(pathname)
	if err != nil {
		return err
	}
	err = db.Update(func(txn *bg.Txn) error {
		lists := Marshal(data)
		e := bg.NewEntry([]byte(key), lists)
		err = txn.SetEntry(e)
		return err
	})
	return err
}

func SetWithTTL(pathname string, key string, data DataConf, ttl int64) error {
	db, err := Open(pathname)
	if err != nil {
		return err
	}
	err = db.Update(func(txn *bg.Txn) error {
		lists := Marshal(data)
		e := bg.NewEntry([]byte(key), lists).WithMeta(0).WithTTL(time.Duration(ttl * time.Second.Nanoseconds()))
		err = txn.SetEntry(e)
		return err
	})
	return err
}
