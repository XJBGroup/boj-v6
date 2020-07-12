package main

import (
	"github.com/Myriad-Dreamin/artisan"
)

type AuthCategories struct {
	artisan.VirtualService
}

func DescribeAuthService() artisan.ProposingService {

	svc := &AuthCategories{}
	svc.Name("AuthService")
	return svc
}
