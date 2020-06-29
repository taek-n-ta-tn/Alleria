package router

import (
	"github.com/taek-n-ta-tn/Alleria/app/pay/demo"
	"github.com/taek-n-ta-tn/Alleria/app/hello"
	"github.com/taek-n-ta-tn/Alleria/app/sys/captcha"
	"github.com/taek-n-ta-tn/Alleria/app/sys/configs"
	"github.com/taek-n-ta-tn/Alleria/app/sys/dept"
	"github.com/taek-n-ta-tn/Alleria/app/sys/dict"
	"github.com/taek-n-ta-tn/Alleria/app/sys/dictData"
	"github.com/taek-n-ta-tn/Alleria/app/sys/menu"
	"github.com/taek-n-ta-tn/Alleria/app/sys/post"
	"github.com/taek-n-ta-tn/Alleria/app/sys/role"
	"github.com/taek-n-ta-tn/Alleria/app/sys/upload"
	"github.com/taek-n-ta-tn/Alleria/app/sys/user"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	//magicToken := r.Header.Get("Authorization")
	//if magicToken == "" {
	//	response.JsonExit(r, common.TOKEN_ERR, common.TOKEN_ERR_MSG)
	//}
	//token, b := util.ParseMagicToken(magicToken)
	//
	//if token == "" {
	//	response.JsonExit(r, common.TOKEN_ERR, common.TOKEN_VALIDE_ERR_MSG)
	//}
	//if !b {
	//	response.JsonExit(r, common.TOKEN_ERR, common.TOKEN_VALIDE_ERR_MSG)
	//}
	r.Middleware.Next()
}

// 用于配置初始化.
func init() {
	s := g.Server()

	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		s.Group("/v1", func(group *ghttp.RouterGroup) {
			group.ALL("/demo", new(demo.Controller))
			group.ALL("/hello", new(hello.Controller))
		})
	})
}
