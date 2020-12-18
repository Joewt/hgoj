package syserror



type Error interface {
	Code() int
	Error() string
	ReasonError() error
}


func New(msg string, reason error) Error {
	return UnKnowError{msg: msg, reason: reason}
}


func NoArt() Error {
	return NoArtError{}
}

func UpdateProErr() Error {
	return ProErr{}
}
