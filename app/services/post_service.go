package services

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/models"
	"lost-and-found-backend/configs/database"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
)

// ================================================== 上传图片 ==================================================

func ValidateImage(file *multipart.FileHeader) (string, error) {
	// 限制图片大小不超过5MB
	const maxSize = 5 * 1024 * 1024
	if file.Size > maxSize {
		return "", errs.ErrFileTooLarge
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
		return "", errs.ErrInvalidFileType
	}

	// 打开文件读取前512字节
	src, err := file.Open()
	if err != nil {
		return "", errs.ErrUploadFailed
	}
	defer src.Close()
	buffer := make([]byte, 512)
	src.Read(buffer)

	// 获取真实的 MIME 类型
	contentType := http.DetectContentType(buffer)
	if !strings.HasPrefix(contentType, "image/") {
		return "", errs.ErrInvalidFileType
	}

	return ext, nil

}

// ================================================== 发布帖子 ==================================================

func CreatePost(userID int, postType, title, eventLocation, eventTime, contact, description, imageUrl string) (*models.Post, error) {
	var post = models.Post{
		UserID:        userID,
		PostType:      postType,
		Title:         title,
		EventLocation: eventLocation,
		EventTime:     eventTime,
		Contact:       contact,
		Description:   description,
		ImageUrl:      imageUrl,
		Status:        "待审核",
	}
	err := database.DB.Model(&models.Post{}).Create(&post).Error
	if err != nil {
		return nil, errs.ErrDatabase
	}

	return &post, nil

}
