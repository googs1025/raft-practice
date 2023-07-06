package pkg

import (
	"github.com/dgraph-io/badger/v2"
	"time"
)

var LocalCache *Bcache

type Bcache struct {
	*badger.DB
}

func NewBcache(path string) *Bcache {
	options := badger.DefaultOptions(path)
	options.Truncate = true //for windows
	db, err := badger.Open(options)
	if err != nil {
		panic(err)
	}
	return &Bcache{DB: db}
}

func (bc *Bcache) SetItem(key string, value string) error {
	err := bc.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	})
	return err
}

//带过期时间的setter
func (bc *Bcache) SetItemWithTTl(key string, value string, ttl time.Duration) error {
	err := bc.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(key), []byte(value)).WithTTL(ttl)
		return txn.SetEntry(e)
	})
	return err
}

func (bc *Bcache) KeysWithPrefix(size int, prefix string) ([]string, error) {
	keys := make([]string, 0)
	err := bc.View(func(txn *badger.Txn) error {
		itopt := badger.IteratorOptions{
			PrefetchValues: false,
			PrefetchSize:   size,
			Reverse:        false,
			AllVersions:    false,
		}
		itor := txn.NewIterator(itopt)
		defer itor.Close()
		prefix_bytes := []byte(prefix)
		for itor.Seek(prefix_bytes); itor.ValidForPrefix(prefix_bytes); itor.Next() {
			key := string(itor.Item().Key())
			keys = append(keys, key)
		}
		return nil
	})
	return keys, err
}

func (bc *Bcache) Keys(size int) ([]string, error) {
	keys := make([]string, 0)
	err := bc.View(func(txn *badger.Txn) error {
		itopt := badger.IteratorOptions{
			PrefetchValues: false,
			PrefetchSize:   size,
			Reverse:        false,
			AllVersions:    false,
		}
		itor := txn.NewIterator(itopt)
		defer itor.Close()
		for itor.Rewind(); itor.Valid(); itor.Next() {
			key := string(itor.Item().Key())
			keys = append(keys, key)
		}
		return nil
	})
	return keys, err
}

func (bc *Bcache) GetItem(key string) (string, error) {
	var ret string
	err := bc.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		_ = item.Value(func(val []byte) error {
			ret = string(val)
			return nil
		})
		return nil
	})
	if err != nil {
		return "", err
	}
	return ret, nil
}
