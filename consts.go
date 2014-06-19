package main

// constants for common usage
const (
	secretKey     = "secret-key"
	sessionKey    = "session-key"
	sessionPrefix = "session_prefix_"
)

// constants for gorilla/sessions.FilesystemStore
const (
	filesystemStorePath = "sessions.txt"
)

// constants for yosssi/boltstore
const (
	boltDBPath = "sessions.db"
)

// constants for bradleypeabody/gorilla-sessions-memcache
const (
	memcacheServer = "localhost:11211"
)

// constants for kidstuff/mongostore
const (
	mongoServer     = "localhost"
	mongoDB         = "test"
	mongoCollection = "test_session"
)

// constants for srinathgs/mysqlstore
const (
	mySQLServer    = "@tcp(localhost:3306)/sessions?parseTime=true&loc=Local"
	mySQLTableName = "sessions"
)

// constants for antonlindstrom/pgstore
const (
	postgreSQLServer = "postgres://%s:%s@127.0.0.1:5432/sessions?sslmode=disable"
)

// constants for boj/redistore
const (
	redisServer = ":6379"
)

// constants for boj/riakstore
const (
	riakServer = "127.0.0.1:8098"
	riakBucket = "sessions"
)
