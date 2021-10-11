package models

import (
	"gorm.io/gorm"
)

// Paste 粘帖结构体
type Paste struct {
	gorm.Model
	Things string `json:"things,omitempty"`     // 粘帖的字符
	Poster string `json:"poster" gorm:"unique"` // 粘帖者 相当于UserName
}

// User 用户结构体
type User struct {
	gorm.Model
	UserName     string `json:"user_name" gorm:"unique"` // 用户名
	EncryptedPwd string `json:"password"`                // 密码
}
