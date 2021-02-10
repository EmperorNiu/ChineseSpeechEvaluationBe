package controlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zhouyongProject/models"
)

func GetSmallWord(c *gin.Context) {
	var words []models.Word
	var word models.Word
	str := c.Query("word")
	t := c.Query("type")
	if err := models.QueryWords(&words,str,t); err != nil {
		c.JSON(http.StatusOK, gin.H{"result":"None"})
	} else {
		if len(words)==0 {
			word.Frequent = 0
			word.Word = str
			word.IsCommon = ""
			words = append(words, word)
		}
		c.JSON(http.StatusOK, gin.H{"result":words})
	}
}