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

// ================================================== 查询帖子 ==================================================

// 查询所有帖子
func GetAllPosts(page int, postType string) ([]models.Post, int, error) {
	var total int64
	var pageSize int = 15

	posts := make([]models.Post, 0)

	query := database.DB.Model(&models.Post{}).Where("status = ?", "已通过")

	//筛选帖子种类
	if postType != "" && (postType == "寻物" || postType == "招领") {
		query = query.Where("post_type = ?", postType)
	}

	//获取帖子总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, errs.ErrDatabase
	}
	offset := (page - 1) * pageSize

	//分页查询
	err = query.Offset(offset).Limit(pageSize).Find(&posts).Error
	if err != nil {
		return nil, 0, errs.ErrDatabase
	}

	return posts, int(total), nil
}

// 查询自己的帖子
func GetMyPosts(userID int, page int, status string, postType string) ([]models.Post, int, error) {
	var total int64
	var pageSize int = 15

	posts := make([]models.Post, 0)

	query := database.DB.Model(&models.Post{}).Where("user_id = ?", userID)

	//筛选帖子种类
	if postType != "" && (postType == "寻物" || postType == "招领") {
		query = query.Where("post_type = ?", postType)
	}

	//筛选审核状态
	if status != "" && (status == "待审核" || status == "已驳回" || status == "已通过") {
		query = query.Where("status = ?", status)
	}

	//获取帖子总数
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, errs.ErrDatabase
	}
	offset := (page - 1) * pageSize

	//分页查询
	err = query.Offset(offset).Limit(pageSize).Find(&posts).Error
	if err != nil {
		return nil, 0, errs.ErrDatabase
	}

	return posts, int(total), nil
}
