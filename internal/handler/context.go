package handler

import (
	"github.com/arielril/hexgo/internal/container"
	uController "github.com/arielril/hexgo/internal/controller/user"
)

type HandlerContext struct {
	UserController uController.UserController
}

func createContainer() *container.Container {
	containerConf := &container.ContainerConfig{
		Database: new(interface{}),
	}

	return container.NewContainer(containerConf)
}

func NewContext() *HandlerContext {
	container := createContainer()

	userCtrl := uController.NewUserController(container)

	return &HandlerContext{
		UserController: userCtrl,
	}
}
