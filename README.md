# Go session store benchmarks

This benchmark suite aims to compare the performance of session stores which implements the [gorilla/sessions](https://github.com/gorilla/sessions) package's [Store](http://godoc.org/github.com/gorilla/sessions#Store) interface in Go.

## Tested session stores

* [CookieStore](http://godoc.org/github.com/gorilla/sessions#CookieStore) - Cookie
* [FilesystemStore](http://godoc.org/github.com/gorilla/sessions#FilesystemStore) - File system
* [yosssi/boltstore](https://github.com/yosssi/boltstore) - Bolt
* [radleypeabody/gorilla-sessions-memcache](https://github.com/bradleypeabody/gorilla-sessions-memcache) - Memcache
* [kidstuff/mongostore](https://github.com/kidstuff/mongostore) - MongoDB
* [srinathgs/mysqlstore](https://github.com/srinathgs/mysqlstore) - MySQL
* [antonlindstrom/pgstore](https://github.com/antonlindstrom/pgstore) - PostgreSQL
* [boj/redistore](https://github.com/boj/redistore) - Redis

## Results

### Get - returns a cached session

```
BenchmarkCookieStore_Get	            20000000	      99.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkFilesystemStore_Get	        20000000	       107 ns/op	       0 B/op	       0 allocs/op
BenchmarkBoltstore_Get	                20000000	       109 ns/op	       0 B/op	       0 allocs/op
BenchmarkGorillaSessionsMemcache_Get	20000000	       109 ns/op	       0 B/op	       0 allocs/op
BenchmarkMongostore_Get	                20000000	       117 ns/op	       0 B/op	       0 allocs/op
BenchmarkMysqlstore_Get	                20000000	       111 ns/op	       0 B/op	       0 allocs/op
BenchmarkPgstore_Get	                20000000	       144 ns/op	       0 B/op	       0 allocs/op
BenchmarkRedistore_Get	                20000000	       130 ns/op	       0 B/op	       0 allocs/op
```

### New - creates and return a new session 

```
BenchmarkCookieStore_New	             5000000	       450 ns/op	     180 B/op	       3 allocs/op
BenchmarkFilesystemStore_New	         5000000	       430 ns/op	     180 B/op	       3 allocs/op
BenchmarkBoltstore_New	                10000000	       287 ns/op	     131 B/op	       2 allocs/op
BenchmarkGorillaSessionsMemcache_New	 5000000	       336 ns/op	     180 B/op	       3 allocs/op
BenchmarkMongostore_New	                 5000000	       395 ns/op	     180 B/op	       3 allocs/op
BenchmarkMysqlstore_New	                 5000000	       427 ns/op	     180 B/op	       3 allocs/op
BenchmarkPgstore_New	                10000000	       301 ns/op	     131 B/op	       2 allocs/op
BenchmarkRedistore_New	                 5000000	       320 ns/op	     131 B/op	       2 allocs/op
```

### Save - persists session to the underlying store implementation

```
BenchmarkCookieStore_Save	              100000	     22740 ns/op	    4101 B/op	      50 allocs/op
BenchmarkFilesystemStore_Save	          10000	    238869 ns/op	    6955 B/op	      81 allocs/op
BenchmarkBoltstore_Save	                    5000	    361275 ns/op	   25184 B/op	     116 allocs/op
BenchmarkGorillaSessionsMemcache_Save	   10000	    129865 ns/op	    7098 B/op	      91 allocs/op
BenchmarkMongostore_Save	                10000	    195949 ns/op	    9938 B/op	     144 allocs/op
BenchmarkMysqlstore_Save	                10000	    342614 ns/op	    6272 B/op	      97 allocs/op
BenchmarkPgstore_Save	                    5000	    577048 ns/op	    9437 B/op	     134 allocs/op
BenchmarkRedistore_Save	                  10000	    150446 ns/op	    5476 B/op	      68 allocs/op
```
