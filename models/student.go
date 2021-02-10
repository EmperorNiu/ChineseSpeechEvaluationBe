package models

import "time"

type Student struct {
	StudentId string `json:"student_id" gorm:"primary_key;"`
	Name string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Teacher string `json:"teacher"`
	KeTangPaiAccount string `json:"ke_tang_pai_account"`
}

type Teacher struct {
	TeacherId int `json:"teacher_id" gorm:"primary_key;auto-increment"`
	Name string `json:"name"`
	Password string `json:"password" gorm:"default:123456"`
	Authority int `json:"authority" gorm:"default:1"`
}

type Password struct {
	Name string `json:"name"`
	Password string `json:"password"`
	NewPassword string `json:"new_password"`
}

type StudentHomework struct {
	StudentHomeworkId int `json:"student_homework_id" gorm:"primary_key;auto-increment"`
	StudentIdRefer string `json:"student_id_refer" gorm:"ForeignKey:StudentId"`
	HomeworkDocIdRefer int `json:"homework_doc_id_refer" gorm:"ForeignKey:HomeworkDocId"`
	CreatedAt time.Time `json:"created_at"`
	Audio string `json:"audio"`
	Type string `json:"type"`
}

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
}

type WordError struct {
	WordErrorId int `json:"word_error_id" gorm:"primary_key"`
	StudendIdRefer string `json:"studend_id_refer" gorm:"ForeignKey:StudentId"`
	StudentHomeworkResultIdRefer int `gorm:"ForeignKey:StudentHomeworkResultId"`
	Word string `json:"word"`
	ErrorTypes string `json:"error_types"`
	WholeWord string `json:"whole_word"`
}

func QueryStudents(stu *[]Student) error{
	if err := db.Find(&stu).Error; err != nil {
		return err
	} else {
		return nil
	}
}
func (s *Student)QueryStudent(id string) error{
	if err := db.Where("student_id = ?",id).First(&s).Error; err != nil {
		return err
	} else {
		return nil
	}
}
func QueryStudentByTeacher(stu *[]Student,teacher string) error{
	if err := db.Where("teacher = ?",teacher).Find(&stu).Error; err != nil {
		return err
	} else {
		return nil
	}
}
func QueryTeachers(t *[]Teacher) error{
	if err := db.Find(&t).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func (s *Student) Insert() error{
	return db.Create(&s).Error
}

func (t *Teacher) Insert() error{
	return db.Create(&t).Error
}

func (t *Teacher) Update(newPassword string) error{
	return db.Model(t).Update("password",newPassword).Error
}

func (t *Teacher) Query(name string) error{
	if err := db.Where("name = ?",name).First(&t).Error; err != nil {
		return err
	} else {
		return nil
	}
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

func QueryAudios(stus *[]StudentHomework,stu_id string,doc_id string) error{
	return db.Where("student_id_refer = ? AND homework_doc_id_refer = ?",stu_id,doc_id).Find(&stus).Error
}

func QueryErrors(errors *[]WordError,stu_id string,result_id string) error{
	return db.Where("studend_id_refer = ? AND student_homework_result_id_refer = ?",stu_id,result_id).Find(&errors).Error
}

func QueryHomeworkResult(results *[]StudentHomeworkResult,stu_id string) error{
	return db.Where("student_id_refer = ?",stu_id).Find(&results).Error
}

func ResultDelete(stu_id string, doc_id string,result_id string) error{
	err := db.Delete(WordError{},"student_homework_result_id_refer = ?", result_id).Error
	//err := db.Delete(StudentHomeworkResult{}, "student_id_refer = ? AND homework_doc_id_refer = ?", stu_id,doc_id).Error
	if err != nil {
		return err
	} else {
		err := db.Delete(StudentHomeworkResult{}, "student_homework_result_id = ?", result_id).Error
		//err := db.Delete(WordError{},"studend_id_refer = ? AND student_homework_result_id_refer = ?", stu_id,result_id).Error
		return err
	}
}

func (result *StudentHomeworkResult) QueryById(result_id string) error{
	if err := db.Where("student_homework_result_id = ?",result_id).First(&result).Error; err != nil {
		return err
	} else {
		return nil
	}
}
