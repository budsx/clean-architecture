package helper

func NewPanicError(err error) {
	if err != nil {
		panic(err)
	}
}
