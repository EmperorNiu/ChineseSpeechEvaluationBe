package controlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	_ "os/exec"
	"strconv"
	"zhouyongProject/e"
	"zhouyongProject/models"
)

func GetWordExercise(c *gin.Context){
	id := c.Query("id")
	var exercises []models.Homework
	if err := models.QueryHomework(id,&exercises);err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"result":exercises})
	}
}

func GetWordExerciseResult(c *gin.Context){
	stu_id := c.Query("stu_id")
	result_id := c.Query("result_id")
	var errors []models.WordError
	if err := models.QueryErrors(&errors,stu_id,result_id);err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"errors": errors})
	}
}

func GetArticle(c *gin.Context){
	id := c.Query("id")
	var art models.HomeWorkArticle
	if err := art.QueryArticle(id);err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"result":art})
	}
}


func UploadDoc(c *gin.Context){
	form,err := c.MultipartForm()
	files := form.File["file"]
	title := form.Value["title"][0]
	des := form.Value["describe"][0]

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	var dst string
	for _,f := range files{
		dst = "./upload/doc/" + f.Filename
		c.SaveUploadedFile(f,dst)
	}
	var homeworkDoc = models.HomeworkDoc{
		Title:         title,
		Describe:      des,
		Position:      dst,
	}
	if err = homeworkDoc.Insert();err != nil{
		c.JSON(http.StatusOK, gin.H{"err":err})
	}
	program := "python"
	arg0 := "./read_homework.py"
	arg2 := strconv.Itoa(homeworkDoc.HomeworkDocId)
	//fmt.Println(dst)
	//fmt.Println(arg2)
	cmd := exec.Command(program,arg0,dst,arg2)
	out, err := cmd.CombinedOutput()
	fmt.Println("concatenation: ", string(out))
	c.JSON(http.StatusOK, gin.H{"homework_id":arg2})
}

func UploadStudentHomework2(c *gin.Context){
	var homeworkDoc models.HomeworkDoc
	if err := c.ShouldBind(&homeworkDoc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS, "message": e.GetMsg(e.INVALID_PARAMS)})
	} else if err := homeworkDoc.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_INSERT,"message":e.GetMsg(e.ERROR_INSERT)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}

func GetDocList(c *gin.Context)  {
	var docs []models.HomeworkDoc
	if err := models.QueryDocList(&docs);err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Error":err})
	} else {
		c.JSON(http.StatusOK, gin.H{"docs":docs})
	}
}

func GetReport(c *gin.Context) {
	program := "python"
	arg0 := "./generateReport5.py"
	arg1 := c.Query("stu_id")
	arg2 := c.Query("doc_id")
	var s models.Student
	var homework models.HomeworkDoc
	s.QueryStudent(arg1)
	homework.QueryHomeWorkDoc(arg2)
	arg3 := s.Name
	arg4 := s.Teacher
	arg5 := homework.Title
	arg6 := homework.CreatedAt.Format("2006.01.02")
	arg7 := ""
	var audios []models.StudentHomework
	models.QueryAudios(&audios,arg1,arg2)
	for _,value := range audios {
		arg7 += value.Audio
	}
	cmd := exec.Command(program,arg0,arg1,arg2,arg3,arg4,arg5,arg6,arg7)
	out, _ := cmd.Output()
	//out1 := string(out)
	//out1 = strings.Replace(out1, " ", "", -1)
	//out1 = strings.Replace(out1, "\n", "", -1)
	fmt.Println(out)
	filename := "./download/doc/report_" + arg1 + "_" + arg2 + ".docx"
	filename2 := "report_" + arg1 + "_" + arg2
	c.Writer.Header().Set("Content-Type", "text/docx")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.docx", filename2))
	http.ServeFile(c.Writer, c.Request, filename)
	//c.File("./download/doc/report_0_5.docx")
}

func GetReport2(c *gin.Context) {
	program := "python"
	arg0 := "./generateReport3.py"
	arg1 := c.Query("stu_id")
	arg2 := c.Query("doc_id")
	var s models.Student
	var homework models.HomeworkDoc
	s.QueryStudent(arg1)
	homework.QueryHomeWorkDoc(arg2)
	arg3 := s.Name
	arg4 := s.Teacher
	arg5 := homework.Title
	arg6 := homework.CreatedAt.Format("2006.01.02")
	arg7 := ""
	var audios []models.StudentHomework
	models.QueryAudios(&audios,arg1,arg2)
	for _,value := range audios {
		arg7 += value.Audio
	}
	cmd := exec.Command(program,arg0,arg1,arg2,arg3,arg4,arg5,arg6,arg7)
	out, _ := cmd.Output()

	fmt.Println(out)
	filename := "./download/doc/report_" + arg1 + "_" + arg2 + ".docx"
	filename2 := "report_" + arg1 + "_" + arg2
	c.Writer.Header().Set("Content-Type", "text/docx")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.docx", filename2))
	http.ServeFile(c.Writer, c.Request, filename)
}

func DeleteHomework(c *gin.Context) {
	doc_id := c.Query("doc_id")
	if err := models.DeleteHomeworkDoc(doc_id); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}