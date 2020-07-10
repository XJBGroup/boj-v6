package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/blog-backend-v2/types"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
)

var codeField = artisan.Param("code", new(types.CodeRawType))

//var required = artisan.Tag("binding", "required")

type Meta struct {
	artisan.RouterMeta
}

func (m *Meta) NeedAuth() *Meta {
	return &Meta{
		RouterMeta: artisan.RouterMeta{
			RuntimeRouterMeta: m.RuntimeRouterMeta,
			NeedAuth:          true,
		},
	}
}

func main() {
	v1 := "v1"

	//instantiate
	userCate := DescribeUserService()
	announcementCate := DescribeAnnouncementService()

	svc := artisan.NewService(
		userCate,
		announcementCate,
	).Base(v1).SetPackageName("api").Final()

	userSvc := svc.GetService(userCate)
	delete(svc.GetServices(), userCate)

	sugar.HandlerError0(svc.PublishRouter("api/router.go"))

	sugar.HandlerError0(svc.PublishObjects(
		userSvc.SetFilePath("api/user.go")))
	sugar.HandlerError0(userSvc.SetFilePath(
		"abstract/control/user-interface.go").PublishInterface(
		"control", svc.Opts))

	announcementSvc := svc.GetService(announcementCate)
	delete(svc.GetServices(), announcementCate)

	sugar.HandlerError0(announcementSvc.SetFilePath(
		"abstract/control/announcement-interface.go").PublishInterface(
		"control", svc.Opts))
	sugar.HandlerError0(svc.PublishObjects(
		announcementSvc.SetFilePath("api/announcement.go")))

	sugar.HandlerError0(svc.Publish())
}
