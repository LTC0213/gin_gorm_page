package Dao

import (
	"gin_gorm_page/Config"
	models "gin_gorm_page/Model"
)

func GetAllUsers(user *models.User, pagination *models.Pagination) (*[]models.User, error) {
	var users []models.User
	// 分页查询
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := Config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&models.User{}).Where(user).Find(&users)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &users, nil
}
