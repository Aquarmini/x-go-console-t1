package controller

import (
	"fmt"
	"app/container"
)

type TestController struct {
	DI container.Container
}

func (this *TestController)Handle() (err error) {
	version, _ := this.DI.Config.GetKey("application", "version")
	r := "Hello World " + version.Value()
	//fmt.Println(this.DI.Config)
	fmt.Println(r)
	return nil
}