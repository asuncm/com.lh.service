package yugabyte

func TABLE(key string) error {
	//db, err := OpenDB("allkic", "")
	//if err != nil {
	//	fmt.Println(err, 1111)
	//	return err
	//}
	//_, err = db.Exec(context.Background(), key)
	//if err != nil {
	//	fmt.Println(err, 22222)
	//	return err
	//}
	//err = db.Close(context.Background())
	//fmt.Println(err, 33333)
	return nil
}

func ADD(key string, value string, keys []string, values interface{}) (interface{}, error) {
	//pool, err := OpenDB(key, value)
	//defer pool.Close()
	//if err != nil {
	//	return nil, err
	//}
	//
	//_, err = pool.CopyFrom(context.Background(),
	//	yugabyte.Identifier{value},
	//	keys,
	//	yugabyte.CopyFromRows(values.([][]any)),
	//)
	//if err != nil {
	//	return nil, err
	//}

	return nil, nil
}
