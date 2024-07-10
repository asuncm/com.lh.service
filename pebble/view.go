package pebble

import (
	"fmt"
	"log"
)

func GetKey(opts Config) (interface{}, interface{}) {
	db, err := OpenDB(opts.Path, opts.Config)
	var msg interface{}
	var data interface{}
	fmt.Println(1, err)
	if err != nil {
		msg = err
	} else {
		val, closer, err := db.Get([]byte(opts.Prefix + "_" + opts.ID))
		fmt.Println(val, closer, err, "-------2")
		if err != nil {
			msg = "get"
		} else {
			data = val
			err = closer.Close()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	err = db.Close()
	if err != nil {
		panic(err)
	}
	return data, msg
}
