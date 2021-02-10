package models

import "time"

type SelectSheet struct {
	SelectSheetId uint `json:"select_sheet_id" gorm:"primary_key" gorm:"auto-increment"`
	CreatedAt time.Time `json:"created_at"`
	Title string `json:"title"`
	Description string `json:"description"`
	Semester int `json:"semester" gorm:"default:1"`
	SelectChars []SelectChar `gorm:"ForeignKey:SelectSheetId"`
	SelectWords []SelectWord `gorm:"ForeignKey:SelectSheetId"`
}

type SelectChar struct {
	SelectCharId uint `json:"select_char_id" gorm:"primary_key;auto-increment" `
	WordDictId int `json:"word_dict_id"`
	Word string `json:"word"`
	Type string `json:"type"`
	StockTimes uint `json:"stock_times" gorm:"default:0;"`
	SelectSheetId uint
	StockId int `gorm:"default:1" json:"stock_id"`
}

type SelectWord struct {
	SelectWordId int `json:"select_word_id" gorm:"primary_key" gorm:"auto-increment"`
	WordDictId int `json:"word_dict_id"`
	Word string `json:"word"`
	IsCommon string `json:"is_common"`
	StockTimes uint `json:"stock_times" gorm:"default:0;"`
	SelectSheetId uint
	StockId int `gorm:"default:1" json:"stock_id"`
}

type SelectStock struct {
	StockWordId uint `json:"stock_word_id" gorm:"primary_key;auto-increment"`
	Semester int `json:"semester"`
	SelectTimes uint `json:"select_times"`
	Content string `json:"content"`
}

func (sheet *SelectSheet) Create() error {
	return db.Create(&sheet).Error
}

func (word *SelectWord) Insert() error {
	return db.Create(&word).Error
}

func (sheet *SelectSheet) Query(id string) error {
	return db.Where("select_sheet_id = ?",id).First(&sheet).Error
}

func QuerySelectSheetList(selectSheets *[]SelectSheet) error {
	if err := db.Find(&selectSheets).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func QuerySheetContent(selectchars *[]SelectChar,selectwords *[]SelectWord,id string) error {
	if err := db.Where("select_sheet_id = ?",id).Find(&selectchars).Error; err != nil {
		return err
	} else if err = db.Where("select_sheet_id = ?",id).Find(&selectwords).Error; err != nil{
		return err
	} else {
		return nil
	}
}

// 注意这里没有做事务处理，后续来改进
func UpdateStock(sheet SelectSheet) error{
	chars := sheet.SelectChars
	words := sheet.SelectWords
	for i:=0;i<len(chars);i++ {
		tmp := SelectStock{}
		if db.Table("select_stock").Where("content = ?",chars[i].Word).First(&tmp).RecordNotFound() {
			tmp = SelectStock{
				Semester:    sheet.Semester,
				SelectTimes: 1,
				Content:     chars[i].Word,
			}
			if err := db.Create(&tmp).Error; err != nil {
				return err
			}
		} else {
			if err:= db.Model(&tmp).Update("select_times", tmp.SelectTimes+1).Error;err!=nil{
				return err
			}
		}
	}
	for i:=0;i<len(words);i++ {
		tmp1 := SelectStock{}
		if db.Table("select_stock").Where("content = ?",words[i].Word).First(&tmp1).RecordNotFound() {
			tmp1 = SelectStock{
				Semester:    sheet.Semester,
				SelectTimes: 1,
				Content:     words[i].Word,
			}
			if err := db.Create(&tmp1).Error; err != nil {
				return err
			}
		} else {
			if err:= db.Model(&tmp1).Update("select_times", tmp1.SelectTimes+1).Error;err!=nil{
				return err
			}
		}
	}
	return nil
}

func QueryHistoryChar(sheets *[]SelectSheet,word string) error {
	chars := []SelectChar{}
	if err := db.Where("word=?",word).Find(&chars).Error; err != nil {
		return err
	} else {
		for i:=0;i<len(chars);i++{
			sheet := SelectSheet{}
			 db.Table("select_sheet").Where("select_sheet_id=?",chars[i].SelectSheetId).First(&sheet)
			*sheets = append(*sheets,sheet)
		}
		return nil
	}
}
func QueryHistoryWord(sheets *[]SelectSheet,word string) error {
	words := []SelectWord{}
	if err := db.Where("word=?",word).Find(&words).Error; err != nil {
		return err
	} else {
		for i:=0;i<len(words);i++{
			sheet := SelectSheet{}
			db.Table("select_sheet").Where("select_sheet_id=?",words[i].SelectSheetId).First(&sheet)
			*sheets = append(*sheets,sheet)
		}
		return nil
	}
}
//func BatchInsertWords(words []*SelectWord) error {
//	var err
//	for _,word := range words {
//		err = db.Create(&word).Error
//		if err != nil{
//			return err
//		}
//	}
//	return nil
//}
//
//func BatchInsertChars(chars []*SelectChar) error {
//	var err
//	for _,char := range chars {
//		err = db.Create(&char)
//		if err != nil{
//			return err
//		}
//	}
//	return nil
//}