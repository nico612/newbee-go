package v1

import "github.com/nico612/newbee-goo/api/v1/manage"

type ApiGroup struct {
	ManageApiGroup manage.ManagerGroup
}

var ApiGroupAPP = new(ApiGroup)
