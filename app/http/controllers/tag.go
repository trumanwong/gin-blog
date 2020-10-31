package controllers

import (
	"gin-blog/app/models"
	"gin-blog/pkg/settings"
	. "gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

type Tag struct {
	Controller
}

func (this *Tag) Get(context *gin.Context) {
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
	res["lists"], res["total"] = tag.Paginate(where, GetPage(context), settings.PageSize)
	Response(context, res, http.StatusOK, "操作成功")
}

//新增文章标签
func (this *Tag) Add(context *gin.Context) {
	tag := models.Tag{}
	tag.Name = context.PostForm("name")
	tag.State = com.StrTo(context.DefaultPostForm("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Required(tag.Name, "name").Message("名称不能为空")
	valid.MaxSize(tag.Name, 100, "name").Message("名称最长为100字符")
	valid.Range(tag.State, 0, 1, "state").Message("状态只允许0或1")

	if valid.HasErrors() {
		Response(context, nil, http.StatusBadRequest, valid.Errors[0].Message)
		return
	}
	tag.GetOneByName(tag.Name)
	if tag.ID > 0 {
		Response(context, nil, http.StatusBadRequest, "该标签已存在!")
		return
	}
	tag.Create()
	Response(context, nil, http.StatusOK, "操作成功")
}

//修改文章标签
func (this *Tag) Update(context *gin.Context) {
	id := com.StrTo(context.Param("id")).MustInt()
	tag := models.Tag{}
	tag.Find(id)

	tag.Name = context.Query("name")
	tag.State = com.StrTo(context.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Required(tag.Name, "name").Message("名称不能为空")
	valid.MaxSize(tag.Name, 100, "name").Message("名称最长为100字符")
	valid.Range(tag.State, 0, 1, "state").Message("状态只允许0或1")

	if valid.HasErrors() {
		Response(context, nil, http.StatusBadRequest, valid.Errors[0].Message)
		return
	}
	tag.Update()
	Response(context, nil, http.StatusOK, "操作成功")
}

//删除文章标签
func (this *Tag) Delete(context *gin.Context) {
	id := com.StrTo(context.Param("id")).MustInt()
	tag := models.Tag{}
	tag.Delete(id)
	Response(context, nil, http.StatusOK, "操作成功")
}
