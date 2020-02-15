package conf

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"gopkg.in/ini.v1"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Server struct {
	Port int
}

type Database struct {
	Driver string `ini:"driver"`
	Username string `ini:"username"`
	Password string `ini:"password"`
	Dbname string `ini:"dbname"`
	Host string `ini:"host"`
	Port int `ini:"port"`
}

var (
	ServerConfig  = &Server{}
	DatabaseConfig = &Database{}
	DB *gorm.DB
)

func init()  {
	//读取配置文件
	cfg, err := ini.Load("./conf/app.ini")
	if err != nil {
		log.Fatal(err)
	}
	port, err := cfg.Section("server").Key("port").Int()
	if err != nil {
		log.Fatal(err)
	}
	ServerConfig.Port = port
	err = cfg.Section("database").MapTo(DatabaseConfig)
	if err != nil {
		log.Fatal(err)
	}
	//连接数据库
	DB, err = gorm.Open(DatabaseConfig.Driver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DatabaseConfig.Username,DatabaseConfig.Password,DatabaseConfig.Host,DatabaseConfig.Port,DatabaseConfig.Dbname))
	if err != nil {
		log.Fatal(err)
	}
	DB.LogMode(true)
	DB.SingularTable(true)
}
