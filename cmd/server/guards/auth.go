package guards

func IsClientAuthorized(token string) bool {
	return len(token) != 0
}
