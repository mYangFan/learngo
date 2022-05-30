package db

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"gonb/pkg/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)


type Mysql struct {
	host string
	port string
	user string
	password string
	database string
	Conn *sql.DB
}

var my *Mysql
var GormDb *gorm.DB
var err error

func New()  *Mysql{
	db := GetDbConfig()
	if my == nil{
		my = &Mysql{
			host:     db["host"],
			port:     db["port"],
			user:     db["user"],
			password: db["password"],
			database: db["database"],
		}

		my.init()
	}

	return my
}

func init(){
	New()
	InitMysqlGorm()
}

func GetDbConfig() map[string]string{
	return viper.GetStringMapString("db")
}

func (m *Mysql) init() {
	my := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.user, m.password, m.host, m.port, m.database)
	m.Conn, _ = sql.Open("mysql", my)
	m.Conn.SetMaxOpenConns(400)
	m.Conn.SetMaxIdleConns(64)
	m.Conn.SetConnMaxLifetime(time.Minute)

	err := m.Conn.Ping()
	if err != nil {
		fmt.Printf("ping mysql error: %+v", err)
	}
}

func InitMysqlGorm()  {
	db := New()
	GormDb, err = gorm.Open(mysql.New(mysql.Config{
		Conn: db.Conn,
	}), &gorm.Config{
		Logger: gormLogger,
	})

	if err != nil {
		log.Printf("init gorm error")
	}
}