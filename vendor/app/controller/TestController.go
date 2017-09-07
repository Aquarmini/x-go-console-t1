package controller

import (
	"fmt"
	"app/container"
	"app/models"
)

type TestController struct {
	DI container.Container
}

func (this *TestController)Handle() (err error) {
	version, _ := this.DI.Config.GetKey("application", "version")
	r := "Hello World " + version.Value()
	fmt.Println(r)

	user := models.User{}
	fmt.Println(user.Find(1))
	return nil
}
