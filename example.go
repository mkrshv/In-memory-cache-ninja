package main

import (
	"fmt"
	"time"

	cache "github.com/mkrshv/In-memory-cache-ninja/internal"
)

func main() {
	c := cache.New()

	c.Set("UserID", 42, time.Second*5)

	if id, err := c.Get("UserID"); err != nil {
		panic(err)
	} else {
		fmt.Println(id)
	}

	time.Sleep(time.Second * 6)

	if id, err := c.Get("UserID"); err != nil {
		panic(err)
	} else {
		fmt.Println(id)
	}
}
