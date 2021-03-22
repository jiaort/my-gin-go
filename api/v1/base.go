package v1

import (
	"my-gin-go/core/response"
	"my-gin-go/global"
	"my-gin-go/model/request"
	resp "my-gin-go/model/response"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

// @Tags Base
// @Summary 生成验证码
// @Accept application/json
// @Produce application/json
// @Success 200 {string} string "{"code":10001,"data":{},"msg":"成功"}"
// @Router /base/captcha [get]
func Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(global.G_CONFIG.Captcha.ImgHeight, global.G_CONFIG.Captcha.ImgWidth, global.G_CONFIG.Captcha.KeyLong, global.G_CONFIG.Captcha.MaxSkew, global.G_CONFIG.Captcha.DotCount)
	cp := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	id, b64s, err := cp.Generate()
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("获取数据失败，%v", err), c)
	} else {
		response.OkWithDataMessage(resp.CaptchaResponse{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功", c)
	}
}

// @Tags Base
// @Summary 验证验证码
// @Accept application/json
// @Produce application/json
// @Param data body request.CaptchaVerify true "验证验证码"
// @Success 200 {string} string "{"code":10001,"data":{},"msg":"成功"}"
// @Router /base/verify [post]
func Verify(c *gin.Context) {
	var params request.CaptchaVerify
	_ = c.ShouldBindJSON(&params)
	fmt.Println(params)
	if base64Captcha.DefaultMemStore.Verify(params.CaptchaId, params.Captcha, true) {
		response.Ok(c)
	} else {
		response.Fail(c)
	}
}
