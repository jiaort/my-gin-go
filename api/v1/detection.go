package v1

import (
	"encoding/json"
	"my-gin-go/core/response"
	"my-gin-go/global"
	resp "my-gin-go/model/response"
	"my-gin-go/service"
	"my-gin-go/utils/upload"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// @Tags Detection
// @Summary 上传文件示例
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Success 200 {string} string "{"code":10001,"data":{},"msg":"上传成功"}"
// @Router /detection/imgDetection [post]
func ImgDetection(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("上传文件失败，%v", err), c)
	} else {
		oss := upload.NewOss()
		filePath, fileName, err := oss.UploadFile(header)
		os.MkdirAll(global.G_CONFIG.Local.Result, os.ModePerm)
		if err != nil {
			response.FailWithMessage(fmt.Sprintf("接收返回值失败，%v", err), c)
		} else {
			args := []string{"python/FirstObjectDetection.py", "box", filePath, path.Join(global.G_CONFIG.Local.Result, fileName)}
			out, err := exec.Command("python3", args...).Output()
			if err != nil {
				response.FailWithMessage(fmt.Sprintf("识别失败，%v", err), c)
			} else {
				// oss.DeleteFile(fileName)
				var result []resp.DetectionResult
				json.Unmarshal(out, &result)

				response.OkWithData(resp.DataResult{
					Result:     result,
					ImgPath:    filePath,
					ResImgPath: "uploads/result/" + fileName,
				}, c)
			}
		}
	}
}

func Ping(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(resquest *http.Request) bool {
			return true
		},
	}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		if string(message) == "ping" {
			out := make(chan string)
			defer close(out)
			go func() {
				for {
					str, ok := <-out
					if !ok {
						break
					}
					ws.WriteMessage(mt, []byte(str))
				}
			}()
			args := []string{"www.baidu.com"}
			if err := service.RunCommand(out, "ping", args...); err != nil {
				return
			}
		}
	}
}
