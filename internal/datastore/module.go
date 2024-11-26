package datastore

type Identifier interface {
	Identity() string
}

type Value = []byte

type Storage interface {
	Store(Identifier, Value) error
	Load(Identifier) (Value, error)
}
