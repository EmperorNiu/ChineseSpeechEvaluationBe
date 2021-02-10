package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"zhouyongProject/controlers"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors())
	auth := router.Group("/api/auth")
	{
		auth.GET("/test", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message":"test",
			})
		})
		//auth.POST("/register",controlers.Register)
		//auth.POST("/login",controlers.Login)
	}
	char := router.Group("/api/char")
	{
		char.GET("/query",controlers.GetChar)
	}
	word := router.Group("/api/word")
	{
		word.GET("/query",controlers.GetSmallWord)
	}
	sheet := router.Group("/api/sheet")
	{
		sheet.POST("/create",controlers.CreatSheet)
		sheet.GET("/getSheet",controlers.GetSheetContent)
		sheet.GET("/getSheetList",controlers.GetSheetList)
		sheet.GET("/getCharHistory",controlers.GetCharHistory)
		sheet.GET("/getWordHistory",controlers.GetWordHistory)

	}
	homework := router.Group("/api/homework")
	{
		homework.POST("/upload/doc",controlers.UploadDoc)
		homework.POST("/upload/review",controlers.UploadStudentHomework2)
		homework.POST("/upload/audio",controlers.UploadDoc)
		homework.GET("/getWordExercise",controlers.GetWordExercise)
		homework.GET("/getWordExerciseResult",controlers.GetWordExerciseResult)
		homework.GET("/getDoc",controlers.GetDocList)
		homework.GET("/getArticle",controlers.GetArticle)
		homework.GET("/getReport",controlers.GetReport)
		homework.GET("/getReport2",controlers.GetReport2)
		homework.GET("/deleteHomework",controlers.DeleteHomework)
	}
	student := router.Group("/api/student")
	{
		student.POST("/upload/audio",controlers.UploadStudentHomework)
		student.POST("/upload/StudentList",controlers.UploadStudentList)
		student.GET("/getAllStudent",controlers.GetStudents)
		student.GET("/getStudentByTeacher",controlers.GetStudentsByTeacher)
		student.POST("/insert",controlers.AddStudent)
		student.POST("/addTeacher",controlers.AddTeacher)
		student.GET("/getTeachers",controlers.GetTeachers)
		student.POST("/homworkresult",controlers.AddHomeworkResult)
		student.GET("/getAudios",controlers.GetStudentAudio)
		student.GET("/getHomeworkResults",controlers.GetStudentHomeworkResults)
		student.GET("/deleteResult",controlers.DeleteHomeworkResult)
		student.GET("/getHomeworkResultScore",controlers.GetHomeworkResultScore)
		student.GET("/getSummary",controlers.GetSummaryReport)
	}
	teacher := router.Group("/api/teacher")
	{
		teacher.POST("/login",controlers.Login)
		teacher.POST("/password",controlers.ChangePassword)
	}
	resource := router.Group("/api/resource")
	{
		resource.GET("/audio", func(context *gin.Context) {
			pos := context.Query("pos")
			context.File(pos)
		})
		resource.GET("studentListTemplate",func(context *gin.Context) {
			filename := "./download/studentList.xlsx"
			filename2 := "studentList"
			context.Writer.Header().Set("Content-Type", "text/xlsx")
			context.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.xlsx", filename2))
			http.ServeFile(context.Writer, context.Request, filename)
		})
	}
	return router
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}