package models

import (
	"errors"
	"time"
)

// 学生信息表
type Student struct {
	StudentId string `json:"student_id" gorm:"primary_key;"`
	Name string `json:"name" gorm:"uniqueIndex"`
	PhoneNumber string `json:"phone_number"`
	Teacher string `json:"teacher"`
	KeTangPaiAccount string `json:"ke_tang_pai_account"`
}

// 学生作业表
type StudentHomework struct {
	StudentHomeworkId int `json:"student_homework_id" gorm:"primary_key;auto-increment"`
	StudentIdRefer string `json:"student_id_refer" gorm:"ForeignKey:StudentId"`
	HomeworkDocIdRefer int `json:"homework_doc_id_refer" gorm:"ForeignKey:HomeworkDocId"`
	CreatedAt time.Time `json:"created_at"`
	Audio string `json:"audio"`
	Type string `json:"type"`
	IsMark int `json:"is_mark" gorm:"default:1"`
}
// 注：is_mark 1: 已批改，2: 未批改

// 学生作业结果表
type StudentHomeworkResult struct {
	StudentHomeworkResultId int `json:"student_homework_result_id" gorm:"primary_key;auto-increment"`
	StudentIdRefer string `json:"student_id_refer" gorm:"ForeignKey:StudentId"`
	HomeworkDocIdRefer int `json:"homework_doc_id_refer" gorm:"ForeignKey:HomeworkDocId"`
	WordErrors []WordError `json:"word_errors" gorm:"ForeignKey:StudentHomeworkResultIdRefer"`
	ToneAccuracy int `json:"tone_accuracy"`
	IntonationAccuracy int `json:"intonation_accuracy"`
	Fluency int `json:"fluency"`
	Comment string `json:"comment"`
	IsThesisExpress int `json:"is_thesis_express" gorm:"default:0"`
	MinusWordErrorScore int `json:"minus_word_error_score"`
	CreatedAt time.Time `json:"created_at"`
}

type WordError struct {
	WordErrorId int `json:"word_error_id" gorm:"primary_key"`
	StudendIdRefer string `json:"studend_id_refer" gorm:"ForeignKey:StudentId"`
	StudentHomeworkResultIdRefer int `gorm:"ForeignKey:StudentHomeworkResultId"`
	Word string `json:"word"`
	ErrorTypes string `json:"error_types"`
	WholeWord string `json:"whole_word"`
}

// 获取所有学生信息
func QueryStudents(stu *[]Student) error{
	if err := db.Find(&stu).Error; err != nil {
		return err
	} else {
		return nil
	}
}

// id查找学生
func (s *Student)QueryStudent(id string) error{
	if err := db.Where("student_id = ?",id).First(&s).Error; err != nil {
		return err
	} else {
		return nil
	}
}

// name查找学生
func (s *Student)QueryStudentByName(name string) error{
	if err := db.Where("name = ?", name).First(&s).Error; err != nil {
		return err
	} else {
		return nil
	}
}

// 获取老师的所有学生
func QueryStudentByTeacher(stu *[]Student,teacher string) error{
	if err := db.Where("teacher = ?",teacher).Find(&stu).Error; err != nil {
		return err
	} else {
		return nil
	}
}

// 获取老师所有学生未判的作业
func QueryUnmarkHomework(students []Student, unmarkedHomework *[]StudentHomework) error {
	for i:=0; i<len(students); i++ {
		var homework []StudentHomework
		if err:=db.Where("student_id_refer = ? AND is_mark = ?", students[i].StudentId, 2).Find(&homework).Error; err != nil {
			return err
		} else {
			*unmarkedHomework = append(*unmarkedHomework, homework...)
		}
	}
	return nil
}

// 插入学生
func (s *Student) Insert() error{
	return db.Create(&s).Error
}

func (sh *StudentHomework) Insert() error{
	return db.Create(&sh).Error
}

func (result StudentHomeworkResult) Insert() error{
	var old StudentHomeworkResult
	db.Where("student_id_refer = ? AND homework_doc_id_refer = ?", result.StudentIdRefer,result.HomeworkDocIdRefer).First(&old)
	db.Delete(StudentHomeworkResult{}, "student_id_refer = ? AND homework_doc_id_refer = ?", old.StudentIdRefer,old.HomeworkDocIdRefer)
	db.Delete(WordError{},"studend_id_refer = ? AND student_homework_result_id_refer = ?", old.StudentIdRefer,old.StudentHomeworkResultId)
	//db.Delete(WordError{},"student_id_refer = ? AND homework_doc_id_refer = ?", result.StudentIdRefer,result.HomeworkDocIdRefer)
	return db.Create(&result).Error
}

func QueryAudios(stus *[]StudentHomework, stuId string, docId string) error{
	return db.Where("student_id_refer = ? AND homework_doc_id_refer = ?", stuId, docId).Find(&stus).Error
}

func QueryErrors(errors *[]WordError, stuId string, resultId string) error{
	return db.Where("studend_id_refer = ? AND student_homework_result_id_refer = ?", stuId, resultId).Find(&errors).Error
}

// 查找某个学生的所有作业结果
func QueryHomeworkResult(results *[]StudentHomeworkResult, stuId string) error{
	return db.Where("student_id_refer = ?", stuId).Find(&results).Error
}

// 查找某个学生的所有作业
func QueryHomeworkByStudent(homework *[]StudentHomework, stuId string) error{
	return db.Where("student_id_refer = ?", stuId).Find(&homework).Error
}

func ResultDelete(stu_id string, doc_id string, resultId string) error{
	err := db.Delete(WordError{},"student_homework_result_id_refer = ?", resultId).Error
	//err := db.Delete(StudentHomeworkResult{}, "student_id_refer = ? AND homework_doc_id_refer = ?", stu_id,doc_id).Error
	if err != nil {
		return err
	} else {
		err := db.Delete(StudentHomeworkResult{}, "student_homework_result_id = ?", resultId).Error
		//err := db.Delete(WordError{},"studend_id_refer = ? AND student_homework_result_id_refer = ?", stu_id,result_id).Error
		return err
	}
}

// 查找某个结果
func (result *StudentHomeworkResult) QueryById(resultId string) error{
	return db.Where("student_homework_result_id = ?", resultId).First(&result).Error
}

// 查找某个学生的某个作业的结果
func (result *StudentHomeworkResult) QueryHomeworkResultByStuDoc(studentId string, homeworkDocId string) error{
	return db.Where("student_id_refer=? " +
		"AND homework_doc_id_refer=?", studentId, homeworkDocId).Find(&result).Error
}

// 更新学生上传作业信息
func UpdateStudentHomework(studentId string, homeworkDocId string, fileName string) error{
	return db.Model(&StudentHomework{}).Where("student_id_refer=? " +
		"AND homework_doc_id_refer=?", studentId, homeworkDocId).Update("audio", fileName).Error
}

// 更新学生老师
func UpdateStudentTeacher(studentIds []string, teacher string) error{
	for i:=0; i<len(studentIds); i++ {
		var student Student
		if db.Where("student_id = ?", studentIds[i]).First(&student); student.Teacher == "" {
			db.Model(&Student{}).Where("student_id = ?", studentIds[i]).Update("teacher", teacher)
		} else {
			return errors.New("ERROR TEACHER")
		}
	}
	return nil
}

