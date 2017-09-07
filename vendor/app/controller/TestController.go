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
	fmt.Println(r)

	fmt.Println(this.DI.DB.Query("SELECT * FROM user WHERE id = ?", 1))
	return nil
}
