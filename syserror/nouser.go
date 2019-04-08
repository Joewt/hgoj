package syserror

type NoUserError struct {
	UnKnowError
}


func (this NoUserError) Code() int{
	return 1002
}

func (this NoUserError) Error() string{
	return "请登录系统"
}
