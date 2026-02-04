package store

// KVStore is a minimal key/value store interface.
type KVStore interface {
	Get(key []byte) []byte
	Set(key, value []byte)
	Delete(key []byte)
}

// CommitStore supports commit/versioning.
type CommitStore interface {
	KVStore
	Commit() (version int64, hash []byte)
	Version() int64
}
