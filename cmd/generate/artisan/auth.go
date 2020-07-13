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
			artisan.Request(),
			artisan.Reply(
				codeField),
		).Method(
			artisan.DELETE, "RemovePolicy",
			artisan.Request(),
			artisan.Reply(
				codeField),
		).Method(
			artisan.GET, "HasPolicy",
			artisan.Request(),
			artisan.Reply(
				codeField),
		),
		GroupingPolicy: artisan.Ink().Path("/policy/group").Method(
			artisan.POST, "AddGroupingPolicy",
			artisan.Request(),
			artisan.Reply(
				codeField),
		).Method(
			artisan.DELETE, "RemoveGroupingPolicy",
			artisan.Request(),
			artisan.Reply(
				codeField),
		).Method(
			artisan.GET, "HasGroupingPolicy",
			artisan.Request(),
			artisan.Reply(
				codeField),
		),
	}
	svc.Name("AuthService")
	return svc
}
