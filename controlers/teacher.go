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
