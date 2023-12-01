package badger

import (
	bg "github.com/dgraph-io/badger/v4"
)

func GetValue(key string, pathname string) (interface{}, error) {
	db, err := Open(pathname)
	if err != nil {
		return nil, err
	}
	var vals []byte
	err = db.View(func(txn *bg.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		vals, err = item.ValueCopy(nil)
		return nil
	})
	if err != nil {
		return nil, err
	} else {
		return Unmarshal(vals), err
	}
}
