package datastore

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"time"

	bolt "go.etcd.io/bbolt"
)

// Upsert ip address in datastore with current time
// incrementing times seen by one if it already exists
func WriteIp(ipAddr net.IP) error {
	if ipAddr == nil {
		log.Fatal("(WriteIp) 'ipAddr' must be value")
	}
	db := OpenDb()
	defer db.Close()
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("ips"))
		if bucket == nil {
			log.Fatal("(WriteIp) Bucket 'ips' isn't found")
		}

		var workRecord IpRecord
		ipRecord := bucket.Get(ipAddr)
		if ipRecord == nil {
			// create new ip record
			workRecord = IpRecord{
				IpAddr:        ipAddr,
				Seen:          1,
				LastUpdatedAt: time.Now(),
				CreatedAt:     time.Now(),
			}
		} else {
			// update current ip record
			err := json.Unmarshal(ipRecord, &workRecord)
			if err != nil {
				log.Fatalf("Error unmarshal ipRecord: %v", err)
			}

			workRecord.Seen = workRecord.Seen + 1
			workRecord.LastUpdatedAt = time.Now()
		}

		putJson, err := json.Marshal(workRecord)
		if err != err {
			log.Fatalf("Error marshalling workRecord: %v", err)
		}

		// Put returns `error`
		return bucket.Put(ipAddr, putJson)
	})

	return err
}

func AddNew(ipAddr net.IP) error {
	db := OpenDb()
	defer db.Close()

	err := db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte("ips"))
		if bucket == nil {
			return errors.New("AddNew: Bucket 'ips' not found")
		}

		workRecord := IpRecord{
			IpAddr:        ipAddr,
			Seen:          1,
			LastUpdatedAt: time.Now(),
			CreatedAt:     time.Now(),
		}

		putJson, err := json.Marshal(workRecord)
		if err != nil {
			return fmt.Errorf("AddNew: Error marshaling new record: %w", err)
		}

		return bucket.Put(ipAddr, putJson)
	})

	return err
}

func IncrementSeen(ipAddr net.IP) (int, error) {
	db := OpenDb()
	defer db.Close()

	seen := 0

	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("ips"))
		if bucket == nil {
			return errors.New("IncrementSeen: Bucket 'ips' not found")
		}
		record := bucket.Get(ipAddr)
		if record == nil {
			return errors.New(fmt.Sprintf("No Record in datastore for ip %s", ipAddr))
		}

		var workRecord IpRecord
		err := json.Unmarshal(record, &workRecord)
		if err != nil {
			return fmt.Errorf("IncrementSeen: Error unmarshaling record: %w", err)
		}

		workRecord.Seen = workRecord.Seen + 1
		workRecord.LastUpdatedAt = time.Now()
		seen = workRecord.Seen

		newRecord, err := json.Marshal(workRecord)
		if err != nil {
			return fmt.Errorf("IncrementSeen: Error marshaling new record: %w", err)
		}

		return bucket.Put(ipAddr, newRecord)
	})

	return seen, err
}

type IpRecord struct {
	IpAddr        net.IP    `json:"ip"`
	Seen          int       `json:"seen"`
	LastUpdatedAt time.Time `json:"timestamp"`
	CreatedAt     time.Time `json:"createdAt"`
}
