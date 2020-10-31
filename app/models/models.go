package models

import (
	"fmt"
	"gin-blog/pkg/settings"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var db *gorm.DB

type Model struct {
	gorm.Model
}

func init() {
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)

	sec, err := settings.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return tablePrefix + defaultTableName;
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}

func (this *Model) BeforeCreate(scope *gorm.Scope) (err error) {
	err = scope.SetColumn("CreatedAt", time.Now())
	err = scope.SetColumn("UpdatedAt", time.Now())
	return
}

func (this *Model) BeforeUpdate(scope *gorm.Scope) (err error) {
	err = scope.SetColumn("UpdatedAt", time.Now())
	return
}