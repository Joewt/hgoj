package syserror

type NoArtError struct {
	UnKnowError
}


func (this NoArtError) Code() int{
	return 1003
}

func (this NoArtError) Error() string{
	return "没有文章"
}


func (this NoArtError)ReasonError() error {
	return this.reason
}
