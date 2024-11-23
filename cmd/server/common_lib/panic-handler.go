package common_lib

func PanicIfError(e error) {
	if e != nil {
		panic(e)
	}
}
