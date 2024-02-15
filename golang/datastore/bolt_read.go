package datastore

import (
	"errors"
	"fmt"
	"log"
	"net"

	bolt "go.etcd.io/bbolt"
)

func IsIpInStore(ip net.IP) bool {
	db := OpenDb()
	defer db.Close()

	var record []byte
	record = nil
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("ips"))
		if bucket == nil {
			return errors.New("IsIpInStore: bucket `ips` doesn't exist")
		}
		record = bucket.Get(ip)
		return nil
	})

	if err != nil {
		log.Fatalf("IsIpInStore error: %s", err)
	}

	return record != nil
}

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
