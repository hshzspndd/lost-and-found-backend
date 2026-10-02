package services

import (
	"lost-and-found-backend/app/errs"
	"lost-and-found-backend/app/models"
	"lost-and-found-backend/configs/database"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"gorm.io/gorm"
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
	err = query.Order("created_at DESC, post_id DESC").Offset(offset).Limit(pageSize).Find(&posts).Error
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
	err = query.Order("created_at DESC, post_id DESC").Offset(offset).Limit(pageSize).Find(&posts).Error
	if err != nil {
		return nil, 0, errs.ErrDatabase
	}

	return posts, int(total), nil
}

// 查询帖子详情
func GetPostDetails(postID int, userID int, role string) (*models.Post, error) {
	var post models.Post
	err := database.DB.Model(&models.Post{}).Where("post_id = ?", postID).First(&post).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.ErrPostNotFound
		}
		return nil, errs.ErrDatabase
	}

	if post.Status != "已通过" {
		// 如果不是“已通过”，只有作者本人和管理员才能看
		isAdmin := role == "系统管理员" || role == "失物招领管理员"
		isAuthor := post.UserID == userID
		if !isAdmin && !isAuthor {
			return nil, errs.ErrPostNotFound
		}
	}
	return &post, nil
}

// ================================================== 删除帖子 ==================================================

// 删除自己的帖子
func DeleteMyPost(PostID, UserID int) error {
	res := database.DB.Model(&models.Post{}).Where("post_id = ? AND user_id = ?", PostID, UserID).Delete(&models.Post{})

	if res.Error != nil {
		return errs.ErrDatabase
	}

	if res.RowsAffected == 0 {
		postExists, err := CheckPostExistByPostID(PostID)
		if err != nil {
			return errs.ErrDatabase
		}
		if !postExists {
			return errs.ErrPostNotFound
		}
		return errs.ErrIsNotYourPost
	}
	return nil
}

func CheckPostExistByPostID(postID int) (bool, error) {
	var post models.Post
	err := database.DB.Model(&models.Post{}).Where("post_id = ?", postID).First(&post).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// ================================================== 管理员查询所有帖子 ==================================================

func AdminGetAllPosts(page int, postType string, status string) ([]models.Post, int, error) {
	var total int64
	var pageSize int = 15

	posts := make([]models.Post, 0)
	query := database.DB.Model(&models.Post{})
	//筛选帖子种类
	if postType != "" && (postType == "寻物" || postType == "招领") {
		query = query.Where("post_type = ?", postType)
	}
	//筛选提子状态
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
	err = query.Order("created_at DESC, post_id DESC").Offset(offset).Limit(pageSize).Find(&posts).Error
	if err != nil {
		return nil, 0, errs.ErrDatabase
	}

	return posts, int(total), nil
}

// ================================================== 审核帖子 ==================================================

// 审核帖子
func AuditPost(postID int, status string) error {
	// 查帖子
	var post models.Post
	err := database.DB.First(&post, postID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errs.ErrPostNotFound // 帖子不存在
		}
		return errs.ErrDatabase // 数据库错误
	}

	if post.Status != "待审核" {
		return errs.ErrStatusInvalid
	}

	// 执行更新
	err = database.DB.Model(&post).Update("status", status).Error
	if err != nil {
		return errs.ErrDatabase
	}

	return nil
}
