package utils

import (
	"server/pkg/utils/role"
)

type MethodOptions struct {
	protected    bool
	allowedRoles map[role.Role]struct{}
}

// If allowedRoles is empty or nil then everyone has access
func NewMethodOptions() *MethodOptions {
	return &MethodOptions{
		protected:    true,
		allowedRoles: nil,
	}
}

func (o *MethodOptions) SetProtected(value bool) *MethodOptions {
	o.protected = value
	return o
}

func (o *MethodOptions) SetAllowedRoles(roles ...role.Role) *MethodOptions {
	allowed := make(map[role.Role]struct{}, len(roles))
	for _, role := range roles {
		allowed[role] = struct{}{}
	}
	o.allowedRoles = allowed
	return o
}

func (o *MethodOptions) IsAllowedForRole(role role.Role) bool {
	if o.allowedRoles == nil || len(o.allowedRoles) == 0 {
		return true
	}
	_, ok := o.allowedRoles[role]
	return ok
}
