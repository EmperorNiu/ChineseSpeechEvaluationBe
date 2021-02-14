package controlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zhouyongProject/e"
	"zhouyongProject/models"
)

func StudentRegister(c *gin.Context) {
	var auth models.StudentAuth
	if err := c.ShouldBind(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS,"message":e.GetMsg(e.INVALID_PARAMS)})
	} else if err := auth.Create(); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success", "username": auth.Username})
	}
}

func StudentLogin(c *gin.Context) {
	var user,_user models.StudentAuth
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"message": err.Error()})
	} else if err := _user.Query(user.Username); err!=nil || _user.Password != user.Password {
		c.JSON(http.StatusBadGateway,gin.H{"status": e.ERROR_PASSWORD, "message": e.GetMsg(e.ERROR_PASSWORD)})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": _user})
	}
}
