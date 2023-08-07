package config

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type Mysql struct {
	Host         string
	Port         int
	User         string
	Password     string
	DBName       string
	Charset      string
	ParseTime    bool
	Loc          string
	MaxIdleConns int
	MaxOpenConns int
}

type Server struct {
	Ip   string
	Port int
}

type Path struct {
	StaticSourcePath string
}

type Config struct {
	Mysql  Mysql
	Server Server
	Path   Path
}

var Global Config

func init() {
	if _, err := toml.DecodeFile("./config/config.ini", &Global); err != nil {
		log.Panicln(err)
	}
}

func DBconnect() string {
	arg := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		Global.Mysql.User, Global.Mysql.Password, Global.Mysql.Host, Global.Mysql.Port,
		Global.Mysql.DBName, Global.Mysql.Charset, Global.Mysql.ParseTime, Global.Mysql.Loc)
	log.Println(arg)
	return arg
}
