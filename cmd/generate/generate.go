package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/blog-backend-v2/types"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
)

var codeField = artisan.Param("code", new(types.CodeRawType))
var required = artisan.Tag("binding", "required")

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
	commentCate := DescribeCommentService()

	svc := artisan.NewService(
		userCate,
		announcementCate,
		commentCate,
	).Base(v1).SetPackageName("api").Final()

	sugar.HandlerError0(svc.PublishRouter("api/router.go"))

	for _, tsk := range []struct {
		cate artisan.ProposingService
		name string
	}{
		{userCate, "user"},
		{announcementCate, "announcement"},
		{commentCate, "comment"},
	} {
		subSvc := svc.GetService(tsk.cate)
		delete(svc.GetServices(), tsk.cate)

		sugar.HandlerError0(svc.PublishObjects(
			subSvc.SetFilePath("api/" + tsk.name + ".go")))
		sugar.HandlerError0(subSvc.SetFilePath(
			"abstract/control/"+tsk.name+"-interface.go").PublishInterface(
			"control", svc.Opts))
	}

	sugar.HandlerError0(svc.Publish())
}
