package config

import (
	"app/controller"
	"app/container"
)

func GetTaskMap(di container.Container) map[string]controller.ControllerInterface {
	TaskMap := make(map[string]controller.ControllerInterface)

	TaskMap["Test"] = &controller.TestController{DI:di}
	TaskMap["Log"] = &controller.LogController{DI:di}

	return TaskMap
}

