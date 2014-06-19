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
