package main

import (
	"errors"
	"fmt"
	"sort"
)

var ErrKeyDoesNotExist = errors.New("Key does not exist")
var ErrEmptyKey = errors.New("Key cannot be empty")
var ErrStoreFull = errors.New("Store is at max capacity")

type Store struct {
	cap  int
	data map[string]string
}

func NewStore(capacity int) *Store {
	return &Store{
		cap:  capacity,
		data: make(map[string]string),
	}
}

func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))

	for key := range s.data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

func (s *Store) Get(key string) (string, error) {
	if key == "" {
		return "", ErrEmptyKey
	}

	val, ok := s.data[key]
	if !ok {
		return "", ErrKeyDoesNotExist
	}
	return val, nil
}

func (s *Store) Set(key, value string) error {
	if key == "" {
		return ErrEmptyKey
	}

	if _, err := s.Get(key); s.cap > 0 && s.Len() >= s.cap && err != nil {
		return fmt.Errorf("Set(%q): %w", key, ErrStoreFull)
	}

	s.data[key] = value

	return nil
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}

func (s *Store) Len() int {
	return len(s.data)
}
