package types

import "time"

//TODO 定义放到单独的文件
type Problem struct {
	ProblemId    int32     `json:"problem_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Input        string    `json:"input"`
	Output       string    `json:"output"`
	SampleInput  string    `json:"sampleinput"`
	SampleOutput string    `json:"sampleoutput"`
	Spj          string    `json:"spj"`
	Hint         string    `json:"hint"`
	Source       string    `json:"source"`
	InDate       time.Time `json:"in_date"`
	TimeLimit    int32     `json:"time_limit"`
	MemoryLimit  int32     `json:"memory_limit"`
	Defunct      string    `json:"defunct"`
	Accepted     int32     `json:"accepted"`
	Submit       int32     `json:"submit"`
	Solved       int32     `json:"solved"`
}
