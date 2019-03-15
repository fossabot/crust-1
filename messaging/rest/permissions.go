package rest

import (
	"context"

	"github.com/crusttech/crust/messaging/internal/service"
	"github.com/crusttech/crust/messaging/rest/request"
)

type Permissions struct {
	svc struct {
		perm service.PermissionsService
	}
}

func (Permissions) New() *Permissions {
	ctrl := &Permissions{}
	ctrl.svc.perm = service.DefaultPermissions
	return ctrl
}

func (ctrl *Permissions) Effective(ctx context.Context, r *request.PermissionsEffective) (interface{}, error) {
	return ctrl.svc.perm.With(ctx).Effective()
}
