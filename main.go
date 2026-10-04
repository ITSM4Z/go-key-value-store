package main

import (
	"fmt"
)

func main() {
	store := NewStore(2)
	_ = store.Set("a", "1")
	_ = store.Set("b", "2")

	if err := store.Set("b", "12"); err != nil {
		fmt.Println(err)
		return
	}

	val, err := store.Get("a")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(val)
}
