package main

import (
	"fmt"
)

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Get(key string) (string, bool) {
	val, ok := s.data[key]
	return val, ok
}

func (s *Store) Set(key, value string) {
	s.data[key] = value
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}

func main() {
	store := NewStore()

	value, exists := store.Get("a")
	fmt.Printf("Test 1: \nValue: %v \nExists: %v\n\n", value, exists)

	store.Set("a", "1")
	value, exists = store.Get("a")
	fmt.Printf("Test 2: \nValue: %v \nExists: %v\n\n", value, exists)

	value, exists = store.Get("b")
	fmt.Printf("Test 3: \nValue: %v \nExists: %v\n\n", value, exists)
}
