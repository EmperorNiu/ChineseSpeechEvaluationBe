package models

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

// 获取所有老师信息
func QueryTeachers(t *[]Teacher) error{
	if err := db.Find(&t).Error; err != nil {
		return err
	} else {
		return nil
	}
}

// 插入老师
func (t *Teacher) Insert() error{
	return db.Create(&t).Error
}

// 更新密码
func (t *Teacher) Update(newPassword string) error{
	return db.Model(t).Update("password",newPassword).Error
}

// name查找老师
func (t *Teacher) Query(name string) error{
	if err := db.Where("name = ?",name).First(&t).Error; err != nil {
		return err
	} else {
		return nil
	}
}