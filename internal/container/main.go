package container

import (
	"github.com/arielril/hexgo/internal/container/database/mysql"
	"github.com/arielril/hexgo/internal/container/user"
)

type Container struct {
	UserService user.UserService
}

type ContainerConfig struct {
}

func NewContainer(c *ContainerConfig) (*Container, error) {
	userRepo, err := mysql.NewUserRepo()

	if err != nil {
		return nil, err
	}

	userSvc := user.NewUserService(userRepo)

	return &Container{
		UserService: userSvc,
	}, nil
}
