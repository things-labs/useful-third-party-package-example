package main

import (
	"fmt"
	"log"

	"go.etcd.io/bbolt"
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bbolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println(db.GoString())

	db.View(func(tx *bbolt.Tx) error {
		var err error

		bk := tx.Bucket([]byte("myBucket"))
		if bk == nil {
			bk, err = tx.CreateBucket([]byte("myBucket"))
			if err != nil {
				return err
			}
		}

		v := bk.Get([]byte("mykey"))
		if v == nil {
			log.Println("v is nil")
		}
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})
}
