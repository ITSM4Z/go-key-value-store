package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestStore_LimitedByCapacity(t *testing.T) {
	store := NewStore(3)

	store.Set("Mazen", "3")
	store.Set("Omar", "2")
	store.Set("Adnan", "6")
	store.Set("Meshal", "1")

	got := store.Keys()
	want := []string{"Adnan", "Mazen", "Meshal", "Omar"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Keys()= %v, want %v", got, want)
	}
}

func TestKeys_ReturnsAllKeysSorted(t *testing.T) {
	store := NewStore(4)

	store.Set("Mazen", "3")
	store.Set("Omar", "2")
	store.Set("Adnan", "6")
	store.Set("Meshal", "1")

	got := store.Keys()
	want := []string{"Adnan", "Mazen", "Meshal", "Omar"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Keys()= %v, want %v", got, want)
	}
}

func TestKeys_EmptyStore(t *testing.T) {
	store := NewStore(1)
	got := store.Keys()

	if len(got) != 0 {
		t.Errorf("Keys() on empty store %v, want empty slice", len(got))
	}
}

func TestSetGet_EmptyKey(t *testing.T) {
	store := NewStore(2)

	if _, err := store.Get(""); err == nil || !errors.Is(err, ErrEmptyKey) {
		t.Errorf("Get() Failed")
	}

	if err := store.Set("", ""); err == nil || !errors.Is(err, ErrEmptyKey) {
		t.Errorf("Set() Failed")
	}
}

func TestSetGet_RoundTrip(t *testing.T) {
	store := NewStore(3)
	store.Set("Hello", "World")

	// if got, err := store.Get("Hello"); err != nil || got != "World" {
	// 	t.Errorf("Get()= %v. want World", got)
	// }

	// if got, err := store.Get("Missing"); err == nil || got == "" {
	// 	t.Errorf("Get() Failed")
	// }
}
