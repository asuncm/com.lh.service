package pebble

import "github.com/cockroachdb/pebble"

func GetKey(opts Config) (DConf, error) {
	db, err := pebble.Open(opts.Path, &pebble.Options{})
	if err != nil {
		return DConf{}, err
	}
	defer db.Close()
	val, closer, err := db.Get([]byte(opts.Key))
	if err != nil {
		return DConf{}, err
	}
	defer closer.Close()
	buf, err := DeCode(string(val))
	if err != nil {
		return DConf{}, err
	}
	return buf, nil
}
