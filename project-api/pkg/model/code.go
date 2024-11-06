package model

import "github.com/Gwen0x4c3/team-sync-server/project-common/errs"

var (
	IllegalMobile = errs.NewError(2001, "手机号格式不正确")
)
