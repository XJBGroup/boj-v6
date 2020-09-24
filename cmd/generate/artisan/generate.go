package main

import (
	"fmt"
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"io/ioutil"
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

// todo move inner service generate
func svcMethods(svc artisan.ServiceDescription) (res string) {
	res = fmt.Sprintf("    %sSignatureXXX() interface{}\n", svc.GetName())
	for _, cat := range svc.GetCategories() {
		res += _svcMethods(cat)
	}
	return
}

func _svcMethods(svc artisan.CategoryDescription) (res string) {
	for _, cat := range svc.GetCategories() {
		res += _svcMethods(cat)
	}
	for _, method := range svc.GetMethods() {
		res += "    " + method.GetName() + "(req *api." + method.GetName() + "Request) (*api." + method.GetName() + "Reply, error) \n"
	}
	return
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

		sugar.HandlerError0(ioutil.WriteFile(
			"abstract/inner-control/"+tsk.name+"-inner-interface.go",
			[]byte(fmt.Sprintf(`
package inner_control

import (
	"github.com/Myriad-Dreamin/boj-v6/api"
)

type Inner%s interface {
%s}`, subSvc.GetName(), svcMethods(subSvc))), 0644))
	}

	sugar.HandlerError0(svc.Publish())
}
