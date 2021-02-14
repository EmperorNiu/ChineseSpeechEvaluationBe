package controlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zhouyongProject/e"
	"zhouyongProject/models"
)

func Login(c *gin.Context) {
	var user,_user models.Teacher
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
	} else if err := _user.Query(user.Name); err!=nil || _user.Password != user.Password {
		c.JSON(http.StatusBadGateway,gin.H{"message":e.GetMsg(e.ERROR_PASSWORD)})
	} else {
		c.JSON(http.StatusOK, gin.H{"user":_user})
	}
}

func ChangePassword(c *gin.Context)  {
	var user models.Teacher
	var pwd models.Password
	if err := c.ShouldBindJSON(&pwd); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
	} else if err := user.Query(pwd.Name); err!=nil || user.Password != pwd.Password {
		c.JSON(http.StatusBadGateway,gin.H{"message":e.GetMsg(e.ERROR_PASSWORD)})
	} else {
		user.Update(pwd.NewPassword)
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}

// 获取老师列表
func GetTeachers(c *gin.Context)  {
	var teachers []models.Teacher
	if err := models.QueryTeachers(&teachers);err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"teachers":teachers})
	}
}

// 添加老师
func AddTeacher(c *gin.Context) {
	var teacher models.Teacher
	if err := c.ShouldBind(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS, "message": e.GetMsg(e.INVALID_PARAMS)})
	} else if err := teacher.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_INSERT,"message":e.GetMsg(e.ERROR_INSERT)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}
