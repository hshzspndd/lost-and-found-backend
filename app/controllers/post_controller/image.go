package post_controller

import (
	"fmt"
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/services"
	"lost-and-found-backend/app/utils"
	"lost-and-found-backend/configs/config"
	"math/rand"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type ImageUrlResp struct {
	ImageUrl string `json:"image_url"`
}

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err == http.ErrMissingFile {
		c.Error(errs.ErrFileNotUploaded)
		c.Abort()
		return
	} else if err != nil {
		c.Error(errs.ErrUploadFailed)
		c.Abort()
		return
	}

	ext, err := services.ValidateImage(file)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	//修改文件名
	newFileName := fmt.Sprintf("%d_%d%s", time.Now().Unix(), rand.Intn(10000), ext)
	dst := filepath.Join("./images", newFileName)

	// 保存文件
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.Error(errs.ErrUploadFailed)
		c.Abort()
		return
	}

	// 拼接返回给前端的 URL
	baseUrl := config.Config.GetString("app.base_url")
	imageUrl := fmt.Sprintf("%s/images/%s", baseUrl, newFileName)

	utils.ResponseSuccess(c, ImageUrlResp{
		ImageUrl: imageUrl,
	})
}
