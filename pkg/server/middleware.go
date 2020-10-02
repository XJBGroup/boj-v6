package server

import (
	ginhelper "github.com/Myriad-Dreamin/boj-v6/app/snippet"
	"github.com/Myriad-Dreamin/boj-v6/config"
	"github.com/Myriad-Dreamin/boj-v6/external"
	"github.com/Myriad-Dreamin/boj-v6/lib/jwt"
	"github.com/Myriad-Dreamin/boj-v6/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"github.com/gin-contrib/cors"
	//"github.com/Myriad-Dreamin/gin-middleware/auth/privileger"
	"strconv"
)

func (srv *Server) PrepareMiddleware() bool {
	srv.jwtMW = jwt.NewMiddleWare(func() *jwt.CustomClaims {
		var cc = new(jwt.CustomClaims)
		cc.CustomField = &types.CustomFields{}
		return cc
	}, func(c controller.MContext, cc *jwt.CustomClaims) error {
		c.Set("uid", strconv.FormatUint(
			uint64(cc.CustomField.(*types.CustomFields).UID), 10))
		return nil
	})
	srv.jwtMW.ExpireSecond = 3600 * 24 * 7
	srv.jwtMW.RefreshSecond = 3600 * 24 * 7

	srv.routerAuthMW = controller.NewMiddleware(srv.Module.RequireImpl(new(*external.Enforcer)).(*external.Enforcer),
		"user:", "uid", ginhelper.MissID, ginhelper.AuthFailed)

	srv.corsMW = cors.New(cors.Config{
		//AllowAllOrigins: true,
		AllowOriginFunc: func(origin string) bool { return true },
		AllowOrigins: []string{
			"http://10.105.242.62:23338", "https://10.105.242.62:23338",
			"http://localhost:23338", "https://localhost:23338"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"X-Total-Count"},
		AllowCredentials: true,
	})

	srv.Module.Provide(config.ModulePath.Middleware.JWT, srv.jwtMW)
	srv.Module.Provide(config.ModulePath.Middleware.RouteAuth, srv.routerAuthMW)
	srv.Module.Provide(config.ModulePath.Middleware.CORS, srv.corsMW)
	return true
}
