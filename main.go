package main

import (
	_ "github.com/taek-n-ta-tn/Alleria/boot"
	_ "github.com/taek-n-ta-tn/Alleria/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
