package controllers

import (
	"fmt"
	"gin-blog/app/models"
	"gin-blog/pkg/settings"
	"gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

type Tag struct {
	Controller
}

func (this *Tag) Get(context *gin.Context)  {
	name := context.Query("name")

	where := make(map[string]interface{})
	res := make(map[string]interface{})

	if name != "" {
		where["name"] = name
	}

	var state int = -1
	if arg := context.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		where["state"] = state
	}
	tag := models.Tag{}
	res["lists"], res["total"] = tag.Paginate(where, util.GetPage(context), settings.PageSize)
	this.Response(context, res, http.StatusOK, "操作成功")
}

//新增文章标签
func (this *Tag) Add(context *gin.Context) {
	params := make(map[string]interface{})
	params["name"] = context.PostForm("name")
	params["state"] = com.StrTo(context.DefaultPostForm("state", "0")).MustInt()

	fmt.Println(params)
	valid := validation.Validation{}
	valid.Required(params["name"], "name").Message("名称不能为空")
	valid.MaxSize(params["name"], 100, "name").Message("名称最长为100字符")
	valid.Range(params["state"], 0, 1, "state").Message("状态只允许0或1")

	tag := models.Tag{}
	if valid.HasErrors() {
		this.Response(context, nil, http.StatusBadRequest, valid.Errors[0].Message)
		return
	}
	if tag.GetOneByName(params["name"].(string)).ID > 0 {
		this.Response(context, nil, http.StatusBadRequest, "该标签已存在!")
		return
	}
	tag.Add(params)
	this.Response(context, nil, http.StatusOK, "操作成功")
}

//修改文章标签
func (this *Tag) Update(c *gin.Context) {
}

//删除文章标签
func (this *Tag) Delete(c *gin.Context) {
}