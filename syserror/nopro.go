package syserror

type NoProError struct {
	UnKnowError
}


func (this NoProError) Code() int{
	return 7002
}

func (this NoProError) Error() string{
	return "没有该题目"
}
