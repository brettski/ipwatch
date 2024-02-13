package main

import (
	"fmt"
	"os"

	cli "github.com/jawher/mow.cli"
)

func main() {
	fmt.Printf("brettski's ipwatch\n\n")

	app := cli.App("go-ipwatch", "brettski's ipwatch (golang edition)")
	app.Version("v version", "0.1.0")
	app.Action = func() { fmt.Println("Use -h for command help.") }
	app.Command("run", "run get against endpoint", cmdRun)
	app.Command("config", "Configuration actions", cmdConfigActions)
	app.Command("data", "Interact with datastore", cmdDataActions)

	app.Run(os.Args)

	// config := getEnvConfig()

	// fmt.Printf("%+v", config)
	// ip := net.ParseIP("192.168.1.1")
	// db := datastore.OpenDb()
	// defer db.Close()

	// err := db.Update(func(tx *bolt.Tx) error {
	// 	bucket, err := tx.CreateBucketIfNotExists([]byte("test"))
	// 	if err != nil {
	// 		return err
	// 	}

	// 	err = bucket.Put(ip, []byte("Some json data"))

	// 	return err
	// })

	// err = db.View(func(tx *bolt.Tx) error {
	// 	bucket := tx.Bucket([]byte("test"))
	// 	value := bucket.Get(ip)
	// 	fmt.Printf("value: %+v\n", string(value))
	// 	return nil
	// })

	// if err != nil {
	// 	panic(err)
	// }
}
