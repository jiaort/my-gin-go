package v1

import (
	"my-gin-go/core/response"
	"my-gin-go/global"
	"my-gin-go/model"
	resp "my-gin-go/model/response"

	"github.com/gin-gonic/gin"
)

// @Tags User
// @Summary 用户登录
// @Accept application/json
// @Produce  application/json
// @Success 200 {string} string "{"code":10001,"data":{},"msg":"成功"}"
// @Router /user/login [get]
func Login(c *gin.Context) {
	response.OkWithMessage("hello 够浪!", c)
}

// @Tags User
// @Summary 用户注册账号
// @Accept application/json
// @Produce  application/json
// @Success 200 {string} string "{"code":10001,"data":{},"msg":"成功"}"
// @Router /user/register [post]
func Register(c *gin.Context) {
	user := &model.User{Username: "mlb", NickName: "dysx", Password: "123456"}
	global.G_DB.Create(&user)
	response.Ok(c)
}

// @Tags User
// @Summary 获取用户列表
// @Accept application/json
// @Produce  application/json
// @Success 200 {string} string "{"code":10001,"data":{},"msg":"成功"}"
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	db := global.G_DB.Model(&model.User{})
	var userList []model.User
	var total int64
	db.Count(&total)
	db.Find(&userList)
	response.OkWithData(resp.PageResult{
		List:  userList,
		Total: total,
	}, c)
}
