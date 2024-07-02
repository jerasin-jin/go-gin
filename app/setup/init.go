package setup

import (
	"github.com/Jerasin/app/controller"
	"github.com/Jerasin/app/repository"
	"github.com/Jerasin/app/service"
)

type Initialization struct {
	userRepo repository.UserRepositoryInterface
	userSvc  service.UserServiceInterface
	UserCtrl controller.UserControllerInterface
}

func NewInitialization(userRepo repository.UserRepositoryInterface,
	userService service.UserServiceInterface,
	userCtrl controller.UserControllerInterface,
) *Initialization {
	return &Initialization{
		userRepo: userRepo,
		userSvc:  userService,
		UserCtrl: userCtrl,
	}
}
