package providers

import "fmt"

type Logger struct {
	id   int64
	name string
}

func NewLogger() *Logger {
	fmt.Println("111")
	return &Logger{
		id:1,
		name:"limx",
	}
}

func (this *Logger)Output() (r int64, err error) {
	this.id++
	r = this.id
	return
}
