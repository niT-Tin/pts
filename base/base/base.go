package base

import (
	"context"
	"errors"
	"github.com/joho/godotenv"
	"golang.org/x/sys/unix"
	"log"
	"os"
	basepb "pasteProject/base/api/gen/go"
)

const (
	tempEnvironmentPath = "/home/liuzehao/GolandProjects/pasteProject/base/conf/envs/db.env"
	devPath             = ""
)

var (
	host     string
	port     string
	user     string
	password string
	db       string
	PreSig   int32 // 上一次base启动时的特征值
	NowSig   int32 // 用于计算base服务启动的次数
)

//func init() {
//	StartSignal = unix.Getpid() + unix.Getppid()
//}

type loadType func(...string) error

func Np(f loadType, file ...string) {
	err := f(file...)
	if err != nil {
		log.Fatalf("load environment variable failed on base.go : %v", err)
	}
}

func Load(path ...string) (DSN string) {
	NowSig = int32(unix.Getpid() + unix.Getppid())
	if len(path) != 0 {
		Np(godotenv.Load, path...)
	} else {
		Np(godotenv.Load, tempEnvironmentPath)
	}
	host = os.Getenv("MYSQL_HOST")
	port = os.Getenv("MYSQL_PORT")
	user = os.Getenv("MYSQL_USERNAME")
	password = os.Getenv("MYSQL_PASSWORD")
	db = os.Getenv("MYSQL_DBNAME")
	return user + ":" + password + "@tcp(" + host + ":" + port + ")" + "/" + db + "?charset=utf8mb4&parseTime=True&loc=Local"
}

type Service struct {
	basepb.UnimplementedBaseServiceServer
}

func (s *Service) GetEnv(context.Context, *basepb.BaseRequest) (*basepb.BaseResponse, error) {
	if host == "" || port == "" || user == "" || password == "" || db == "" {
		return &basepb.BaseResponse{}, errors.New("load env error")
	}
	return &basepb.BaseResponse{
		Mysql: &basepb.MySql{
			Username: user,
			Host:     host,
			Port:     port,
			Password: password,
			Dbname:   db,
			Dsn:      user + ":" + password + "@tcp(" + host + ":" + port + ")" + "/" + db + "?charset=utf8mb4&parseTime=True&loc=Local",
		},
		Mongo: &basepb.Mongo{},
		Redis: &basepb.Redis{},
		Sig:   NowSig,
	}, nil
}
