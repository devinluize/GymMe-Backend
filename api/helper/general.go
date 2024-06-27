package helper

func Paniciferror(err error) {
	if err != nil {
		panic(err)
	} else {
		return
	}
}
