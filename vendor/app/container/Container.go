package container

import "app/providers"

type Container struct {
	Logger *providers.Logger `inject:"logger"`
	Config *providers.Config `inject:"config"`
}
