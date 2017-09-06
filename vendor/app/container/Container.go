package container

import "fmt"
import "app/providers"

type Container struct {
	Logger *providers.Logger   `inject:""`
	Config *providers.Config `inject:""`
}

func (this *Container) Render() {
	fmt.Println(this.Logger.Output())
}
