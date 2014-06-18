package main

import (
	"net/http"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/yosssi/boltstore/store"
)

func BenchmarkBoltstore_Get(b *testing.B) {
	// Open a Bolt database.
	db, err := bolt.Open(boltDBPath, 0666)
	if err != nil {
		b.Error(err)
	}

	defer db.Close()

	// Create a store.
	str, err := store.New(db, store.Config{}, []byte(secretKey))
	if err != nil {
		b.Error(err)
	}

	// Create a request.
	req, err := http.NewRequest("GET", "http://localhost:3000/", nil)
	if err != nil {
		b.Error(err)
	}

	// Get a session.
	_, err = str.Get(req, sessionKey)
	if err != nil {
		b.Error(err)
	}
}
