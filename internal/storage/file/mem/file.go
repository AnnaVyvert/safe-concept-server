package mem

// type FileStorage struct {
// 	Logger *slog.Logger
// }

// type ID struct {
// 	user, app int
// }

// var files map[ID][]byte = make(map[ID][]byte)

// func (f *FileStorage) DeleteFile(user int, app int) ([]byte, error) {
// 	const op = "MemFileStorage.DeleteFile"
// 	log := f.Logger.With(slog.String("op", op))
// 	data, err := f.FindFile(user, app)
// 	if err != nil {
// 		log.Warn("not found", sl.Err(err))
// 		return data, err
// 	}
// 	delete(files, ID{user, app})
// 	log.Info("ok")
// 	return data, nil
// }

// func (f *FileStorage) FindFile(user int, app int) ([]byte, error) {
// 	const op = "MemFileStorage.FindFile"
// 	log := f.Logger.With(slog.String("op", op))
// 	data, ok := files[ID{user, app}]
// 	if !ok {
// 		log.Warn("not found")
// 		return nil, fmt.Errorf("%s: %w", op, errors.New("not found"))
// 	}
// 	log.Info("ok", slog.String("data", string(data)))
// 	return data, nil
// }

// func (f *FileStorage) SaveFile(user int, app int, data []byte) error {
// 	const op = "MemFileStorage.SaveFile"
// 	log := f.Logger.With(slog.String("op", op))
// 	log.Info("ok", slog.String("data", string(data)))
// 	// TODO(mxd): mxd return old
// 	files[ID{user, app}] = data
// 	return nil

// }
