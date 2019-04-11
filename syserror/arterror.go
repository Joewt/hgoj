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

type ProErr struct {
	UnKnowError
}


func (this ProErr) Code() int {
	return 1004
}

func (this ProErr) Error() string {
	return "问题数据保存失败"
}

func (this ProErr) ReasonError() error {
	return this.reason
}
