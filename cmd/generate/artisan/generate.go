package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
)

var codeField = artisan.Param("code", new(types.ServiceCode))
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
	var meta = []struct {
		cate artisan.ProposingService
		name string
	}{
		{DescribeUserService(), "user"},
		{DescribeAuthService(), "auth"},
		{DescribeAnnouncementService(), "announcement"},
		{DescribeCommentService(), "comment"},
		{DescribeSubmissionService(), "submission"},
		{DescribeProblemService(), "problem"},
		{DescribeContestService(), "contest"},
		{DescribeGroupService(), "group"},
	}

	var svc *artisan.PublishedServices
	var svcs []artisan.ProposingService
	for _, tsk := range meta {
		svcs = append(svcs, tsk.cate)
	}
	svc = artisan.NewService(svcs...).Base(v1).SetPackageName("api").Final()

	sugar.HandlerError0(svc.PublishRouter("api/router.go"))

	for _, tsk := range meta {
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
