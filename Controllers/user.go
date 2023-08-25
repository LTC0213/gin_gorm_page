package Controllers

import (
	dao "gin_gorm_page/Dao"
	models "gin_gorm_page/Model"
	utils "gin_gorm_page/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetAllUsers(c *gin.Context) {
	pagination := utils.GeneratePaginationFromRequest(c)
	var user models.User
	userLists, err := dao.GetAllUsers(&user, &pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"data": userLists,
	})

}
