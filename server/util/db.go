package util

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var G_db *gorm.DB

func openDB(url string, maxIdleConns, maxOpenConns int, enableLog bool) (err error) {
	if G_db, err = gorm.Open("mysql", url); err != nil {
		return err
	}
	//设置数据库参数
	G_db.LogMode(enableLog)
	G_db.SingularTable(true)
	G_db.DB().SetMaxIdleConns(maxIdleConns)
	G_db.DB().SetMaxOpenConns(maxOpenConns)
	return nil
}

func InitDB() error {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		G_conf.MySqlConfig.UserName, G_conf.MySqlConfig.Password, G_conf.MySqlConfig.IP, G_conf.MySqlConfig.Port, G_conf.MySqlConfig.DataBase)
	err := openDB(url, G_conf.MySqlConfig.MaxIdleConns, G_conf.MySqlConfig.MaxOpenConns, true)
	if err != nil {
		return err
	}
	return nil
}
