package main

import (
	"github.com/Myriad-Dreamin/artisan"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
)

var codeField = artisan.Param("code", new(types.ServiceCode))
var required = artisan.Tag("binding", "required")
var routeParam = artisan.Tag("route-param", "-")

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

//func controllerMethods(controller artisan.ServiceDescription) (res string) {
//	res = fmt.Sprintf("    %sSignatureXXX() interface{}\n", controller.GetName())
//	for _, cat := range controller.GetCategories() {
//		res += _controllerMethods(cat)
//	}
//	return
//}

//func _controllerMethods(controller artisan.CategoryDescription) (res string) {
//	for _, cat := range controller.GetCategories() {
//		res += _controllerMethods(cat)
//	}
//	for _, method := range controller.GetMethods() {
//		res += "    " + method.GetName() + "(c controller.MContext, req *api." + method.GetName() + "Request) (*api." + method.GetName() + "Reply, error) \n"
//	}
//	return
//}

func main() {
	v1 := "v1"

	//instantiate
	var meta = []struct {
		cate artisan.ProposingService
		name string
	}{
		{DescribeUserController(), "user"},
		{DescribeAuthController(), "auth"},
		{DescribeAnnouncementController(), "announcement"},
		{DescribeCommentController(), "comment"},
		{DescribeSubmissionController(), "submission"},
		{DescribeProblemController(), "problem"},
		{DescribeContestController(), "contest"},
		{DescribeGroupController(), "group"},
	}

	var controller *artisan.PublishedServices
	var controllers []artisan.ProposingService
	for _, tsk := range meta {
		controllers = append(controllers, tsk.cate)
	}
	controller = artisan.NewService(controllers...).Base(v1).SetPackageName("api").Final()

	sugar.HandlerError0(controller.PublishRouter("api/router.go"))

	for _, tsk := range meta {
		subController := controller.GetService(tsk.cate)
		delete(controller.GetServices(), tsk.cate)

		sugar.HandlerError0(controller.PublishObjects(
			subController.SetFilePath("api/" + tsk.name + ".go")))
		sugar.HandlerError0(subController.SetFilePath(
			"abstract/control/"+tsk.name+"-interface.go").PublishInterface(
			"control", controller.Opts))

		//		sugar.HandlerError0(ioutil.WriteFile(
		//			"abstract/inner-control/"+tsk.name+"-inner-interface.go",
		//			[]byte(fmt.Sprintf(`
		//package inner_control
		//
		//import (
		//	"github.com/Myriad-Dreamin/boj-v6/api"
		//    "github.com/Myriad-Dreamin/minimum-lib/controller"
		//)
		//
		//type Inner%s interface {
		//%s}`, subController.GetName(), controllerMethods(subController))), 0644))
	}

	sugar.HandlerError0(controller.Publish())
}
