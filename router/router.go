package router

import (
	"github.com/gogf/gf-demos/app/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	// 分组路由注册方式
	s.Group("/", func(group *ghttp.RouterGroup) {
		//group.Middleware(
		//	service.Middleware.Ctx,
		//	service.Middleware.CORS,
		//)

		//宠物列表
		group.ALL("/pet", api.PetApi)

		//group.ALL("/pet", func(r *ghttp.Request) {
		//	r.Response.Write("test")
		//})

		//group.ALL("/chat", api.Chat)
		//group.ALL("/user", api.User)
		//group.ALL("/getInfo", api.Test)
		//
		//group.Group("/", func(group *ghttp.RouterGroup) {
		//	group.Middleware(service.Middleware.Auth)
		//	group.ALL("/user/profile", api.User.Profile)
		//})
	})

}
