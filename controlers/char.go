package controlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zhouyongProject/models"
)

func GetChar(c *gin.Context) {
	var char models.Char
	str := c.Query("char")
	t := c.Query("type")
	if err := char.Query(str,t); err != nil {
		char.Frequent = 0
		char.Word = str
		c.JSON(http.StatusOK, gin.H{"result":char})
	} else {
		c.JSON(http.StatusOK, gin.H{"result":char})
	}
}