package main

import (
	"net/http"
	"testing"

	"github.com/boltdb/bolt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/bradleypeabody/gorilla-sessions-memcache"
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

	// Reset the timer.
	b.ResetTimer()

	// Get a session.
	for i := 0; i < b.N; i++ {
		_, err = str.Get(req, sessionKey)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGorillaSessionsMemcache(b *testing.B) {
	// Create a client.
	client := memcache.New(memcacheServer)

	// Create a store.
	str := gsm.NewMemcacheStore(client, sessionPrefix, []byte(secretKey))

	// Create a request.
	req, err := http.NewRequest("GET", "http://localhost:3000/", nil)
	if err != nil {
		b.Error(err)
	}

	// Reset the timer.
	b.ResetTimer()

	// Get a session.
	for i := 0; i < b.N; i++ {
		_, err = str.Get(req, sessionKey)
		if err != nil {
			b.Error(err)
		}
	}

}
