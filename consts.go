package main

// constants for common usage
const (
	secretKey  = "secret-key"
	sessionKey = "session-key"
)

// constants for yosssi/boltstore
const (
	boltDBPath = "sessions.db"
)

// constants for bradleypeabody/gorilla-sessions-memcache
const (
	memcacheServer = "localhost:11211"
	sessionPrefix  = "session_prefix_"
)
