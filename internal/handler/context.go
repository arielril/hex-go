package handler

import (
	"fmt"

	"github.com/arielril/hexgo/internal/container"
	user "github.com/arielril/hexgo/internal/controller/user"
)

type HandlerContext struct {
	UserController user.UserController
}

func createContainer() *container.Container {
	containerConf := &container.ContainerConfig{}

	cont, err := container.NewContainer(containerConf)

	if err != nil {
		fmt.Println("failed container", err)
		panic("failed to create container")
	}

	return cont
}

func NewContext() *HandlerContext {
	container := createContainer()

	userCtrl := user.NewUserController(container)

	return &HandlerContext{
		UserController: userCtrl,
	}
}
