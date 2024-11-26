package get_file

func getFileNameByToken(token string) string {
	// find client session by token and hash his login
	return "mock_file_name"
}

func decryptContent(fileContent []byte) []byte {
	// decrypt file content
	return fileContent
}

func getClientFile(token string) (string, []byte) {
	clientFileName := getFileNameByToken(token)

	return clientFileName, decryptContent(getFileFromFS(clientFileName))
}
