package core

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

func InitConn() *gorm.DB {
	switch Conf.DB.DBType {
	case "mysql":
		return GormMysql()
	case "sqlite":
		return GormSqlite()
	default:
		log.Fatal("No DBType")
		return nil
	}
}

func initTables(db *gorm.DB) {
	//checkMigrate(&model.BookMark{}, db)
	//checkMigrate(&model.Site{}, db)
	//checkMigrate(&model.WebScreen{}, db)
}

func checkMigrate(tb interface{}, db *gorm.DB) {
	if db.HasTable(tb) == false {
		db.AutoMigrate(tb)
	}
}

func GormMysql() *gorm.DB {
	m := Conf.DB
	if m.DBName == "" {
		return nil
	}
	dsn := m.DBUser + ":" + m.DBPwd + "@tcp(" + m.DBHost + ")/" + m.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("MySQL启动异常", err.Error())
		os.Exit(0)
	}
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(300)
	db.SingularTable(true)
	db.LogMode(true)
	initTables(db)
	return db
}

func GormSqlite() *gorm.DB {
	dbFile := fmt.Sprintf("%s.db", Conf.DB.DBName)
	db, err := gorm.Open("sqlite3", dbFile)
	if err != nil {
		if err := createDB(dbFile); err != nil {
			log.Fatal(err)
		}
	}
	db.SingularTable(true)
	db.LogMode(true)
	initTables(db)
	return db
}

func createDB(path string) error {
	fp, err := os.Create(path) // 如果文件已存在，会将文件清空。
	if err != nil {
		return err
	}
	defer fp.Close() //关闭文件，释放资源。
	return nil
}
