package controller

import (
	"app/container"
	log "github.com/sirupsen/logrus"
	"fmt"
)

type LogController struct {
	DI container.Container
}

func (this *LogController)Handle() (err error) {
	version, _ := this.DI.Config.GetKey("application", "version")
	log.Info(fmt.Sprintf("Current:version=%s", version.Value()));
	fmt.Println("查看日志")
	return nil
}
