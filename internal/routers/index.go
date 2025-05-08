package routers

import (
	"github.com/sondoannam/go-ecommerce-backend-api/internal/routers/manager"
	"github.com/sondoannam/go-ecommerce-backend-api/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manager.ManagerRouterGroup
}

var RouterGroupApp = new(RouterGroup)