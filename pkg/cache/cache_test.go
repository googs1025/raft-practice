package cache

import (
	"fmt"
	"log"
	"testing"
)

func TestCache(t *testing.T) {
	c := NewMapCache()

	err := c.SetItem("test", "test-res")
	if err != nil {
		log.Fatal("set err: ", err)
	}
	res, err := c.GetItem("test")
	if err != nil {
		log.Fatal("get err: ", err)
	}
	fmt.Println(res)
}
