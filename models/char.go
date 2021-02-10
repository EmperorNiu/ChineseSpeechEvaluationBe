package models

type Char struct {
	Index int `json:"index"`
	Word string `json:"word"`
	Frequent int `json:"frequent"`
	Type string `json:"type"`
	Pinyin string
	Shengmu string
	Yunmu1 string
	Yunmu2 string
	Zhuyuanyin string
	Tone int `json:"tone"`
}

type CharStock struct {
	BigCharStorckId int `json:"index"`
	Word string `json:"character"`
	Type string `json:"type"`
	Pinyin string
	Shengmu string
	Yunmu1 string
	Yunmu2 string
	Zhuyuanyin string
	Tone int `json:"tone"`
}


func (charactor *Char) Query(char string,t string) error {
	if(t == "1") {
		return db.Table("smallcharstock").Where("word = ?",char).First(&charactor).Error
	} else {
		return db.Table("middlecharstock").Where("word = ?",char).First(&charactor).Error
	}
}

//func (charactor *Char) Query(char string) error {
//	return db.Table("smallcharstock").Where("character = ?",char).First(&charactor).Error
//}
