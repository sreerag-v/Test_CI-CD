package handler

import (
	"Test/Test-Crud/database"
	"Test/Test-Crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(404, gin.H{"msg": err.Error()})
		c.Abort()
		return
	}
	validationErr := validate.Struct(user)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
		return
	}
	result2 := database.InitDB().Create(&user)
	if result2.Error != nil {
		c.JSON(500, gin.H{
			"Status": "False",
			"Error":  result2.Error.Error(),
		})
	}
	c.JSON(200, user)
}

func GetUserById(c *gin.Context) {
	Uid := c.Param("id")

	DB := database.InitDB()

	var body models.User

	err := DB.Table("users").Where("id = ?", Uid).Scan(&body).Error
	if err != nil {
		c.JSON(500, gin.H{
			"Status": "False",
			"Error":  err,
		})
	}
	c.JSON(200, body)
}

func UpdateUserByid(c *gin.Context) {
	Uid := c.Param("id")
	DB := database.InitDB()

	type userStruct struct {
		First_Name string
		Last_Name  string
		Email      string
		Password   string
	}
	var user userStruct

	if c.Bind(&user) != nil {
		c.JSON(400, gin.H{
			"error": "Data binding error",
		})
		return
	}
	var body models.User
	err := DB.Table("users").
		Where("id = ?", Uid).
		Updates(map[string]interface{}{
			"first_name": user.First_Name,
			"last_name":  user.Last_Name,
			"email":      user.Email,
			"password":   user.Password,
		}).Scan(&body).Error

	if err != nil {
		c.JSON(500, gin.H{
			"Status": "False",
			"Error":  err,
		})
	}
	c.JSON(200, body)
}
