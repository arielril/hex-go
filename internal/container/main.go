package container

import (
	"github.com/arielril/hexgo/internal/container/user"
)

type Container struct {
	UserService user.UserService
}

type ContainerConfig struct {
	Database interface{}
}

func NewContainer(c *ContainerConfig) *Container {
	userRepo := user.NewUserRepo(c.Database)
	userSvc := user.NewUserService(userRepo)

	return &Container{
		UserService: userSvc,
	}
}
