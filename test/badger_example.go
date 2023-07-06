package main

import (
	"github.com/dgraph-io/badger/v2"
	"log"
)

func main() {
	db, err := badger.Open(
		badger.DefaultOptions("tmp"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	key := []byte("name")

	// 写操作
	value := []byte("jiang")
	tx := db.NewTransaction(true)
	if err := tx.Set(key, value); err != nil {
		log.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	// 读操作
	tx2 := db.NewTransaction(false)
	defer tx.Discard()
	item, err := tx2.Get(key)
	if err != nil {
		log.Fatal(err)
	}
	_ = item.Value(func(val []byte) error {
		log.Println(string(val))
		return nil
	})

}
