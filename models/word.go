package models

type Word struct {
	Index int `json:"index"`
	Word string `json:"word"`
	Frequent int `json:"frequent"`
	Pinyin string `json:"pinyin"`
	IsCommon string `json:"is_common"`
}

type WordStock struct {
	BigWordStockId int `json:"index"`
	IsCommon string `json:"is_common"`
	Word string `json:"word"`
	Pinyin string `json:"pinyin"`
}

func QueryWords(words *[]Word,char string,t string) error {
	if(t == "1") {
		return db.Table("smallwordstock").Where("word like ? ","%"+char+"%").Find(&words).Error
	} else {
		return db.Table("middlewordstock").Where("word like ? ","%"+char+"%").Find(&words).Error
	}
}