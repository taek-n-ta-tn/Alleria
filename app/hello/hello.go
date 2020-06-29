package hello

import (
	"github.com/taek-n-ta-tn/Alleria/app"
	"github.com/taek-n-ta-tn/Alleria/lib/response"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
	app.Controller
}

func (c *Controller) Index(r *ghttp.Request) {
	response.SuccessResult(r, "hello word")
}
