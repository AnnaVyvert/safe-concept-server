package response

type FileResponse struct {
	Data  string `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func FileOK() FileResponse {
	return FileResponse{}
}
func FileData(data []byte) FileResponse {
	return FileResponse{
		Data: string(data),
	}
}

func FileError(msg string) FileResponse {
	return FileResponse{
		Error: msg,
	}
}
