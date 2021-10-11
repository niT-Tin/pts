package errs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"pasteProject/base"
	"pasteProject/models"
	"time"
)

var (
	DB *gorm.DB
)

// 此处初始化数据库
func init() {
	var err error
	DB, err = gorm.Open(mysql.Open(base.Load()), &gorm.Config{})
	if err != nil {
		log.Fatalf("init mysql falied: %v", err)
	}
	e := NewErrs(DB)
	err = DB.AutoMigrate(&Err{})
	if err != nil {
		log.Fatalf("auto migrate err failed: %v", err)
	}
	// 初始化User数据表
	errUser := DB.AutoMigrate(&models.User{})

	if errUser != nil {
		e.ReciteErrors(Err{
			//ErrBasic: errUser,
			Message: "初始化User数据表错误",
			When:    time.Now(),
			Where:   "loggers.go",
		})
	}
	// 初始化Paste数据表
	errPaste := DB.AutoMigrate(&models.Paste{})
	if errPaste != nil {
		e.ReciteErrors(Err{
			//ErrBasic: errPaste,
			Message: "初始化Paste数据表错误",
			When:    time.Now(),
			Where:   "loggers.go",
		})
	}

}

func GetDB() *gorm.DB {
	return DB
}

// Err 错误结构体
type Err struct {
	gorm.Model
	//ErrBasic error
	Message string    `json:"error_message"`
	When    time.Time `json:"when"`
	Where   string    `json:"where"`
}

type Errs struct {
	db *gorm.DB
}

type IErr interface {
	ReciteErrors(Err)
	GetAllErrorMsg() []Err
}

func NewErrs(db *gorm.DB) IErr {
	return &Errs{
		db: db,
	}
}

// ReciteErrors 记录错误
func (es *Errs) ReciteErrors(E Err) {
	es.db.Create(E)
}

// GetAllErrorMsg 获取所有错误
func (es *Errs) GetAllErrorMsg() []Err {
	var ES []Err
	es.db.Find(&ES)
	return ES
}
