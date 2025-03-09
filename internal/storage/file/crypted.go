package file

var _ Storage = crypted{}

type crypted struct {
	Storage
}

func Crypted(storage Storage) Storage {
	return &crypted{storage}
}

// FIXME(mxd): impl
func (c crypted) Load(id FileID) (Value, error) {
	data, err := c.Storage.Load(id)
	// decrypt data
	return data, err
}

// FIXME(mxd): impl
// Store implements Storage.
func (c crypted) Store(id FileID, data Value) error {

	// data = crypt(data)
	return c.Storage.Store(id, data)
}
