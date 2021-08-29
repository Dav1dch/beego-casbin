package routers

import (
	"encoding/json"
	"log"

	"github.com/beego/beego/v2/server/web/context"
)

func TokenFilter(ctx *context.Context) {
	method := ctx.Request.Method
	if method == "POST" {
		var requestBody map[string]interface{}
		json.Unmarshal(ctx.Input.RequestBody, &requestBody)
		log.Println(requestBody)
		if requestBody["token"] == nil {
			log.Println("deny")
			denyRequest(ctx, "no token")
		}
	}

}
