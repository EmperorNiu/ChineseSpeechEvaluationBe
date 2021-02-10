package controlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zhouyongProject/e"
	"zhouyongProject/models"
)


// 创建新的sheet，并更新selectStock
func CreatSheet(c *gin.Context) {
	var sheet models.SelectSheet
	if err := c.ShouldBind(&sheet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS, "message": e.GetMsg(e.INVALID_PARAMS)})
	} else if err := sheet.Create(); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_INSERT,"message":e.GetMsg(e.ERROR_INSERT)})
	} else {
		models.UpdateStock(sheet)
		c.JSON(http.StatusOK, gin.H{"message": sheet})
	}
}

func GetSheetContent(c *gin.Context) {
	selectSheetId := c.Query("select_sheet_id")
	var chars []models.SelectChar
	var words []models.SelectWord
	if err := models.QuerySheetContent(&chars,&words,selectSheetId); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"chars":chars,"words":words})
	}
}

func GetSheetList(c *gin.Context) {
	var sheets []models.SelectSheet
	if err := models.QuerySelectSheetList(&sheets);err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"sheets":sheets})
	}
}

func GetCharHistory(c *gin.Context){
	var sheets []models.SelectSheet
	word := c.Query("word")
	if err := models.QueryHistoryChar(&sheets,word);err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"sheets":sheets})
	}
}

func GetWordHistory(c *gin.Context){
	var sheets []models.SelectSheet
	word := c.Query("word")
	if err := models.QueryHistoryWord(&sheets,word);err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"sheets":sheets})
	}
}