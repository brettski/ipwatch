package datastore

import (
	"fmt"

	bolt "go.etcd.io/bbolt"
)

// Dumps all datastore records to stdout
func DumpAllRecords() {
	db := OpenDb()
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("ips"))
		cursor := bucket.Cursor()

		fmt.Println("Dumping all data:")
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			fmt.Printf("[%v] => %s\n", k, v)
		}

		return nil
	})
}
