package middlewares

import (
	"fmt"
	"gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type Jwt struct {
}

func (this *Jwt) Handle() gin.HandlerFunc  {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		token = strings.Split(token, " ")[1]
		claims, err := util.ParseToken(token)
		fmt.Println(claims["expire_at"])
		if err != nil || time.Now().Format("2006-01-02 15:04:05") > claims["expire_at"].(string) {
			util.Response(context, nil, http.StatusForbidden, "操作失败")
			context.Abort()
			return
		}
		context.Next()
	}
}