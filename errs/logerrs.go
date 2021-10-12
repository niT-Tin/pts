package errs

import (
	"context"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	basepb "pasteProject/base/api/gen/go"
	"pasteProject/base/base"
	"pasteProject/models"
	"time"
)

var (
	DB     *gorm.DB // 数据库实例
	DSN    string   // 当前DSN
	conn   *grpc.ClientConn
	client basepb.BaseServiceClient
)

// 启动客户端获取mysql配置信息
func loadMysql() string {
	var err error
	conn, err = grpc.Dial("localhost:8084", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("load: %v", err)
	}
	client = basepb.NewBaseServiceClient(conn)
	res, err := client.GetEnv(context.Background(), &basepb.BaseRequest{Code: 1})
	if err != nil {
		log.Fatalf("load get env: %v", err)
	}
	if DSN == "" {
		base.PreSig = res.Sig
		base.NowSig = res.Sig
	}
	//fmt.Println(res.Mysql.Dsn)
	return res.Mysql.Dsn
}

// InitMysql 初始化数据库
func InitMysql(DSN string) {
	//DSN = loadMysql()
	var err error
	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
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

// 此处初始化数据库
func init() {
	InitMysql(loadMysql())
}

func Refresh() {
	env, err := client.GetEnv(context.Background(), &basepb.BaseRequest{Code: 1})
	if err != nil && base.PreSig == 0 {
		log.Fatalf("get env failed: %v", err)
	} else if err != nil && base.PreSig != 0 {
		return
	}
	if base.PreSig != env.Sig {
		InitMysql(loadMysql())
		base.PreSig = env.Sig
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
