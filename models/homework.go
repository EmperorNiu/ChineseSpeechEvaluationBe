package models

import "time"

type Homework struct {
	Index int `json:"index"`
	Character string `json:"character"`
	Word1 string `json:"word1"`
	Word2 string `json:"word2"`
	HomeworkIndex int `json:"homework_index"`
}

type HomeworkDoc struct {
	HomeworkDocId int `json:"homework_doc_id" gorm:"primary_key;auto-increment"`
	Title string `json:"title"`
	Describe string `json:"describe"`
	Position string `json:"position"`
	CreatedAt time.Time `json:"created_at"`
}

type HomeWorkArticle struct {
	ArticleId int `json:"article_id"`
	Content string `json:"content"`
	HomeworkIndex int `json:"homework_index"`
}


func (doc *HomeworkDoc)Insert() error{
	return db.Create(&doc).Error
}
func QueryHomework(homework_id string,exercises *[]Homework) error{
	if err:=db.Table("homework").Where("homework_index = ?",homework_id).Find(&exercises).Error; err!=nil{
		return err
	} else {
		return nil
	}
}
func (art *HomeWorkArticle)QueryArticle(homework_id string)  error{
	if err:=db.Table("homework_article").Where("homework_index = ?",homework_id).Find(&art).Error; err!=nil{
		return err
	} else {
		return nil
	}
}

func QueryDocList(docs *[]HomeworkDoc) error{
	if err := db.Find(&docs).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func (doc *HomeworkDoc)QueryHomeWorkDoc(id string) error {
	if err := db.Where("homework_doc_id = ?",id).First(&doc).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func DeleteHomeworkDoc(id string)  error{
	err := db.Delete(HomeworkDoc{}, "homework_doc_id = ?",id).Error
	if err != nil {
		return err
	} else {
		err := db.Delete(Homework{},"homework_index = ?", id).Error
		return err
	}
}