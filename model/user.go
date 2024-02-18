package model

type User struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Bato_point int    `json:"bato_point"`
	Safe       bool   `json:"safe"`
}
