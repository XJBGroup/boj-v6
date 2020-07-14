package main

import (
	"github.com/Myriad-Dreamin/artisan"
)

type AuthCategories struct {
	artisan.VirtualService

	Policy         artisan.Category
	GroupingPolicy artisan.Category
}

func DescribeAuthService() artisan.ProposingService {

	svc := &AuthCategories{
		Policy: artisan.Ink().Path("/policy").Method(
			artisan.POST, "AddPolicy",
			artisan.Request(
				artisan.Param("subject", artisan.String, required),
				artisan.Param("object", artisan.String, required),
				artisan.Param("action", artisan.String, required),
			),
			artisan.Reply(
				codeField,
				artisan.Param("data", artisan.Bool)),
		).Method(
			artisan.DELETE, "RemovePolicy",
			artisan.Request(
				artisan.Param("subject", artisan.String, required),
				artisan.Param("object", artisan.String, required),
				artisan.Param("action", artisan.String, required)),
			artisan.Reply(
				codeField,
				artisan.Param("data", artisan.Bool)),
		).Method(
			artisan.GET, "HasPolicy",
			artisan.Request(
				artisan.Param("subject", artisan.String, required),
				artisan.Param("object", artisan.String, required),
				artisan.Param("action", artisan.String, required)),
			artisan.Reply(
				codeField,
				artisan.Param("data", artisan.Bool)),
		),
		GroupingPolicy: artisan.Ink().Path("/policy/group").Method(
			artisan.POST, "AddGroupingPolicy",
			artisan.Request(
				artisan.Param("subject", artisan.String, required),
				artisan.Param("group", artisan.String, required)),
			artisan.Reply(
				codeField,
				artisan.Param("data", artisan.Bool)),
		).Method(
			artisan.DELETE, "RemoveGroupingPolicy",
			artisan.Request(
				artisan.Param("subject", artisan.String, required),
				artisan.Param("group", artisan.String, required)),
			artisan.Reply(
				codeField,
				artisan.Param("data", artisan.Bool)),
		).Method(
			artisan.GET, "HasGroupingPolicy",
			artisan.Request(
				artisan.Param("subject", artisan.String, required),
				artisan.Param("group", artisan.String, required)),
			artisan.Reply(
				codeField,
				artisan.Param("data", artisan.Bool)),
		),
	}
	svc.Name("AuthService")
	return svc
}
