package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"labix.org/v2/mgo"

	"github.com/antonlindstrom/pgstore"
	"github.com/boj/redistore"
	"github.com/boltdb/bolt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/bradleypeabody/gorilla-sessions-memcache"
	"github.com/gorilla/sessions"
	"github.com/kidstuff/mongostore"
	"github.com/srinathgs/mysqlstore"
	"github.com/yosssi/boltstore/store"
)

func BenchmarkCookieStore_New(b *testing.B) {
	// Create a store.
	str := sessions.NewCookieStore([]byte(secretKey))

	// Benchmark.
	benchmarkNew(b, str)
}

func BenchmarkFilesystemStore_New(b *testing.B) {
	// Create a store.
	str := sessions.NewFilesystemStore(filesystemStorePath, []byte(secretKey))

	// Benchmark.
	benchmarkNew(b, str)
}

func BenchmarkBoltstore_New(b *testing.B) {
	// Open a Bolt database.
	db, err := bolt.Open(boltDBPath, 0666, nil)
	if err != nil {
		b.Error(err)
	}

	defer db.Close()

	// Create a store.
	str, err := store.New(db, store.Config{}, []byte(secretKey))
	if err != nil {
		b.Error(err)
	}

	// Benchmark.
	benchmarkNew(b, str)
}

func BenchmarkGorillaSessionsMemcache_New(b *testing.B) {
	// Create a client.
	client := memcache.New(memcacheServer)

	// Create a store.
	str := gsm.NewMemcacheStore(client, sessionPrefix, []byte(secretKey))

	// Benchmark.
	benchmarkNew(b, str)
}

func BenchmarkMongostore_New(b *testing.B) {
	// Create a database session.
	dbsess, err := mgo.Dial(mongoServer)
	if err != nil {
		panic(err)
	}

	defer dbsess.Close()

	// Create a store.
	str := mongostore.NewMongoStore(dbsess.DB(mongoDB).C(mongoCollection), 3600, true, []byte(secretKey))

	// Benchmark.
	benchmarkNew(b, str)
}

func BenchmarkMysqlstore_New(b *testing.B) {
	// Create a store.
	str, err := mysqlstore.NewMySQLStore(os.Getenv("MYSQL_USER")+":"+os.Getenv("MYSQL_PASSWORD")+mySQLServer, mySQLTableName, "/", 3600, []byte(secretKey))
	if err != nil {
		panic(err)
	}

	defer str.Close()

	// Benchmark.
	benchmarkNew(b, str)
}

func BenchmarkPgstore_New(b *testing.B) {
	// Create a store.
	str := pgstore.NewPGStore(fmt.Sprintf(postgreSQLServer, os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD")), []byte(secretKey))

	defer str.Close()

	// Benchmark.
	benchmarkNew(b, str)
}

func BenchmarkRedistore_New(b *testing.B) {
	// Create a store.
	str, err := redistore.NewRediStore(10, "tcp", redisServer, "", []byte(secretKey))
	if err != nil {
		panic(err)
	}

	defer str.Close()

	// Benchmark.
	benchmarkNew(b, str)
}

/*
func BenchmarkRiakstore_New(b *testing.B) {
	// Create a store.
	str := riakstore.NewRiakStore([]string{riakServer}, 5, riakBucket, []byte(secretKey))

	defer str.Close()

	// Benchmark.
	benchmarkNew(b, str)
}
*/

func benchmarkNew(b *testing.B, str sessions.Store) {
	// Create a request.
	req, err := http.NewRequest("GET", "http://localhost:3000/", nil)
	if err != nil {
		b.Error(err)
	}

	// Reset the timer.
	b.ResetTimer()

	// Get a session.
	for i := 0; i < b.N; i++ {
		_, err = str.New(req, sessionKey)
		if err != nil {
			b.Error(err)
		}
	}

	// Stop the timer.
	b.StopTimer()
}
