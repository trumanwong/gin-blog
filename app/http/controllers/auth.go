package controllers

import (
	"fmt"
	"gin-blog/app/models"
	. "gin-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthController struct {
	Controller
}

func (this *AuthController) Login(context *gin.Context)  {
	username := strings.TrimSpace(context.DefaultPostForm("username", ""))
	password := strings.TrimSpace(context.DefaultPostForm("password", ""))
	if len(username) == 0 || len(password) == 0 {
		Response(context, nil, http.StatusBadRequest, "请输入用户名或密码")
		return
	}
	user := &models.User{}
	user, err := user.CheckAuth(username, password)
	if err != nil {
		Response(context, nil, http.StatusBadRequest, err.Error())
		return
	}

	res := make(map[string]interface{})
	res["token"], err = GenerateToken(username, password)
	if err != nil {
		fmt.Println(err.Error())
		Response(context, nil, http.StatusBadRequest, "操作失败")
		return
	}
	Response(context, res, http.StatusOK, "操作成功")
}