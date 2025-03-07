package file

type FileID struct {
	App, User int
}

func ID(app, user int) FileID {
	return FileID{app, user}
}

type Value = []byte

type Storage interface {
	Store(FileID, Value) error
	Load(FileID) (Value, error)
	Delete(FileID) error
}
