package badger

func Delete(key string, pathname string) error {
	db, err := Open(pathname)
	if err != nil {
		return err
	}
	wb := db.NewWriteBatch()
	defer wb.Cancel()
	return wb.Delete([]byte(key))
}
