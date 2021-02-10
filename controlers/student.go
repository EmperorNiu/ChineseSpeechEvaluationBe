package controlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"strconv"
	"zhouyongProject/e"
	"zhouyongProject/models"
)

func GetStudents(c *gin.Context)  {
	var students []models.Student
	if err := models.QueryStudents(&students);err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"students":students})
	}
}

func GetStudentsByTeacher(c *gin.Context)  {
	var students []models.Student
	var teacher = c.Query("teacher")
	if err := models.QueryStudentByTeacher(&students,teacher);err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"students":students})
	}
}

func GetTeachers(c *gin.Context)  {
	var teachers []models.Teacher
	if err := models.QueryTeachers(&teachers);err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"teachers":teachers})
	}
}

func AddStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBind(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS, "message": e.GetMsg(e.INVALID_PARAMS)})
	} else if err := student.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_INSERT,"message":e.GetMsg(e.ERROR_INSERT)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}

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

func AddHomeworkResult(c *gin.Context){
	var homeworkResult models.StudentHomeworkResult
	if err := c.ShouldBind(&homeworkResult); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS, "message": e.GetMsg(e.INVALID_PARAMS)})
	} else if err := homeworkResult.Insert(); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_INSERT,"message":e.GetMsg(e.ERROR_INSERT)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}

func GetStudentAudio(c *gin.Context){
	stu_id := c.Query("stu_id")
	doc_id := c.Query("doc_id")
	var sh []models.StudentHomework
	if err := models.QueryAudios(&sh,stu_id,doc_id); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"students_homework":sh})
	}
}

func GetStudentHomeworkResults(c *gin.Context){
	stu_id := c.Query("stu_id")
	var results []models.StudentHomeworkResult
	if err := models.QueryHomeworkResult(&results,stu_id); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Error": err})
	} else {
		var homeworks []models.HomeworkDoc
		for i:=0;i<len(results);i++ {
			var homeworkdoc models.HomeworkDoc
			doc_id := strconv.Itoa(results[i].HomeworkDocIdRefer)
			homeworkdoc.QueryHomeWorkDoc(doc_id)
			homeworks = append(homeworks, homeworkdoc)
		}
		c.JSON(http.StatusOK, gin.H{"results":results,"homework":homeworks})
	}
}

func DeleteHomeworkResult(c *gin.Context) {
	stu_id := c.Query("stu_id")
	doc_id := c.Query("doc_id")
	result_id := c.Query("result_id")
	if err := models.ResultDelete(stu_id,doc_id,result_id); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}

// 根据 result_id 获取此次评判的分数结果
func GetHomeworkResultScore(c *gin.Context) {
	result_id := c.Query("result_id")
	var result models.StudentHomeworkResult
	if err := result.QueryById(result_id); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"Error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": result})
	}
}

// 获取学生总结报告
func GetSummaryReport(c *gin.Context) {
	program := "python"
	arg0 := "./summary.py"
	arg1 := c.Query("stu_id")
	arg2 := c.Query("stu_name")
	cmd := exec.Command(program,arg0,arg1,arg2)
	out, _ := cmd.Output()

	fmt.Println(out)
	filename := "D:/summary_report/" + arg1 + ".docx"
	filename2 := "summary_report_" + arg1
	c.Writer.Header().Set("Content-Type", "text/docx")
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.docx", filename2))
	http.ServeFile(c.Writer, c.Request, filename)
	//c.File("./download/doc/report_0_5.docx")
}

// 上传学生列表，并调用脚本更改数据库
func UploadStudentList(c *gin.Context) {
	form,err := c.MultipartForm()
	files := form.File["file"]
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	var dst string
	for _,f := range files{
		dst = "./upload/student/" + f.Filename
		c.SaveUploadedFile(f,dst)
	}
	program := "python"
	arg0 := "./insertStudent.py"
	cmd := exec.Command(program,arg0,dst)
	out, err := cmd.CombinedOutput()
	fmt.Println("concatenation: ", string(out))
	c.JSON(http.StatusOK, gin.H{"filename":dst})
}

// 上传学生作业
func UploadStudentHomework(c *gin.Context) {
	form,err := c.MultipartForm()
	files := form.File["file"]
	student_id := form.Value["student_id"][0]
	doc_id,_ := strconv.Atoi(form.Value["doc_id"][0])
	t := form.Value["type"][0]
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	for _,f := range files{
		dst := "D:/upload/audio/" + f.Filename
		c.SaveUploadedFile(f,dst)
		var studentHomework = models.StudentHomework {
			StudentIdRefer:         student_id,
			HomeworkDocIdRefer:      doc_id,
			Audio:      			dst,
			Type:					t,
		}
		if err = studentHomework.Insert(); err != nil {
			c.JSON(http.StatusOK, gin.H{"err":err})
		}
	}
}