package main

import (
	"github.com/Myriad-Dreamin/artisan/artisan-core"
)

type AuthCategories struct {
	artisan_core.VirtualService

	Policy         artisan_core.Category
	GroupingPolicy artisan_core.Category
}

func DescribeAuthController() artisan_core.ProposingService {

	controller := &AuthCategories{
		Policy: artisan_core.Ink().Path("/policy").Method(
			artisan_core.POST, "AddPolicy",
			artisan_core.Request(
				artisan_core.Param("subject", artisan_core.String, required),
				artisan_core.Param("object", artisan_core.String, required),
				artisan_core.Param("action", artisan_core.String, required),
			),
			artisan_core.Reply(
				codeField,
				artisan_core.Param("data", artisan_core.Bool)),
		).Method(
			artisan_core.DELETE, "RemovePolicy",
			artisan_core.Request(
				artisan_core.Param("subject", artisan_core.String, required),
				artisan_core.Param("object", artisan_core.String, required),
				artisan_core.Param("action", artisan_core.String, required)),
			artisan_core.Reply(
				codeField,
				artisan_core.Param("data", artisan_core.Bool)),
		).Method(
			artisan_core.GET, "HasPolicy",
			artisan_core.Request(
				artisan_core.Param("subject", artisan_core.String, required),
				artisan_core.Param("object", artisan_core.String, required),
				artisan_core.Param("action", artisan_core.String, required)),
			artisan_core.Reply(
				codeField,
				artisan_core.Param("data", artisan_core.Bool)),
		),
		GroupingPolicy: artisan_core.Ink().Path("/policy/group").Method(
			artisan_core.POST, "AddGroupingPolicy",
			artisan_core.Request(
				artisan_core.Param("subject", artisan_core.String, required),
				artisan_core.Param("group", artisan_core.String, required)),
			artisan_core.Reply(
				codeField,
				artisan_core.Param("data", artisan_core.Bool)),
		).Method(
			artisan_core.DELETE, "RemoveGroupingPolicy",
			artisan_core.Request(
				artisan_core.Param("subject", artisan_core.String, required),
				artisan_core.Param("group", artisan_core.String, required)),
			artisan_core.Reply(
				codeField,
				artisan_core.Param("data", artisan_core.Bool)),
		).Method(
			artisan_core.GET, "HasGroupingPolicy",
			artisan_core.Request(
				artisan_core.Param("subject", artisan_core.String, required),
				artisan_core.Param("group", artisan_core.String, required)),
			artisan_core.Reply(
				codeField,
				artisan_core.Param("data", artisan_core.Bool)),
		),
	}
	controller.Name("AuthController")
	return controller
}
