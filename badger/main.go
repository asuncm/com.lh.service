package badger

import (
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/badger/v4"
	"os"
)

func Open(pathname string) (*badger.DB, error) {
	_, err := os.Stat(pathname)
	if os.IsNotExist(err) {
		os.MkdirAll(pathname, 0755)
	}
	opts := badger.DefaultOptions(pathname)
	db, err := badger.Open(opts)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	return db, err
}

type DataConf = map[string]interface{}

func Marshal(data DataConf) []byte {
	list := make(DataConf)
	for key, val := range data {
		list[key] = val
		fmt.Println(key, "key", val)
	}
	bytes, _ := json.Marshal(list)
	return bytes
}

func Unmarshal(data []byte) DataConf {
	list := make(DataConf)
	err := json.Unmarshal(data, &list)
	if err != nil {
		return DataConf{}
	}
	return list
}
