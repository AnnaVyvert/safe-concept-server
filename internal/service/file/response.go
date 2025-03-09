package file

type FileResponse struct {
	Data  jsonValue `json:"data,omitempty"`
	Error string    `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func FileOK() FileResponse {
	return FileResponse{}
}
func FileData(data jsonValue) FileResponse {
	return FileResponse{
		Data: data,
	}
}

func FileError(msg string) FileResponse {
	return FileResponse{
		Error: msg,
	}
}
