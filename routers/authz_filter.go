package routers

import (
	"beego-casbin/authz"
	"beego-casbin/controllers"
	"beego-casbin/util"
	"log"

	"github.com/beego/beego/v2/server/web/context"
)

func denyRequest(ctx *context.Context, message string) {
	w := ctx.ResponseWriter
	w.WriteHeader(403)
	// resp := &controllers.Response{Status: "error", Msg: "Unauthorized operation"}
	resp := &controllers.Response{Status: "error", Msg: message}
	_, err := w.Write([]byte(util.StructToJson(resp)))
	if err != nil {
		panic(err)
	}
}

func AuthzFilter(ctx *context.Context) {
	user := "admin"
	method := ctx.Request.Method
	urlPath := ctx.Request.URL.Path

	isAllowed := authz.IsAllowed(user, method, urlPath)

	result := "deny"
	if isAllowed {
		result = "allow"
	}

	log.Println(result)
	if !isAllowed {
		denyRequest(ctx, "not allowd")
	}
}
