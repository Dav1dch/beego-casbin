package main

import (
	"beego-casbin/authz"
	"beego-casbin/routers"
	_ "beego-casbin/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	authz.InitAuthz()
	beego.InsertFilter("*", beego.BeforeRouter, routers.TokenFilter)
	beego.InsertFilter("*", beego.BeforeRouter, routers.AuthzFilter)
	beego.InsertFilter("/signup", beego.BeforeRouter, routers.AccountBloomFilter)
	beego.Run()
}
