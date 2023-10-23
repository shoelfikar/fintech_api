package pkg


func PanicIfError(err interface{}) {
	if err != nil {
		panic(err)
	}
}