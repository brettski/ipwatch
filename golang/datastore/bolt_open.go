package datastore

import (
	"log"

	bolt "go.etcd.io/bbolt"
)

func OpenDb() *bolt.DB {
	db, err := bolt.Open("boltdb.data", 0600, nil)
	if err != nil {
		log.Fatalf("Error opening bolt db: %v", err)
	}
	//defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("ips"))
		return err
	})

	if err != nil {
		log.Fatalf("Error opening db & bucket: %v", err)
	}

	return db
}
