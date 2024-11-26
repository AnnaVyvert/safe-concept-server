package datastore

var _ Storage = crypted{}

type crypted struct {
	inner Storage
}

func Crypted(storage Storage) Storage {
	return &crypted{inner: storage}
}

// Load implements Storage.
func (c crypted) Load(id Identifier) ([]byte, error) {
	data, err := c.inner.Load(id)
	// decrypt data
	return data, err
}

// Store implements Storage.
func (c crypted) Store(id Identifier, data []byte) error {
	// data = crypt(data)
	return c.inner.Store(id, data)
}
