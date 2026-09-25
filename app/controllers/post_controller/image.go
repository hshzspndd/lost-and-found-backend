package post_controller

import (
	"fmt"
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/utils"
	"math/rand"
	"net/http"
	"path/filepath"
	"strings"
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

	filename := file.Filename     //提取文件名
	ext := filepath.Ext(filename) //提取扩展名
	ext = strings.ToLower(ext)    //转成小写

	//限制图片格式
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
	}

	if !allowedExts[ext] {
		c.Error(errs.ErrInvalidFileType)
		c.Abort()
		return
	}

	//修改文件名
	newFileName := fmt.Sprintf("%d_%d%s", time.Now().Unix(), rand.Intn(1000), ext)

	dst := filepath.Join("./images", newFileName)

	// 保存文件
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.Error(errs.ErrUploadFailed)
		c.Abort()
		return
	}

	// 拼接返回给前端的 URL
	imageUrl := fmt.Sprintf("http://localhost:8080/images/%s", newFileName)

	utils.ResponseSuccess(c, ImageUrlResp{
		ImageUrl: imageUrl,
	})
}
