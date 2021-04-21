package core

import (
	"fmt"
	"github.com/cilidm/base-system-v2/app/global"
	"github.com/cilidm/base-system-v2/app/model"
	"github.com/gchaincl/dotsql"
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

var (
	db  *gorm.DB
	err error
)

func GormMysql() *gorm.DB {
	m := Conf.DB
	if m.DBName == "" {
		return nil
	}
	dsn := m.DBUser + ":" + m.DBPwd + "@tcp(" + m.DBHost + ")/" + m.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("MySQL启动异常", err.Error())
		os.Exit(0)
	}
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(300)
	db.SingularTable(true)
	db.LogMode(true)
	initTables()
	initMyCallbacks()
	return db
}

func GormSqlite() *gorm.DB {
	dbFile := fmt.Sprintf("%s.db", Conf.DB.DBName)
	db, err = gorm.Open("sqlite3", dbFile)
	if err != nil {
		if err := createDB(dbFile); err != nil {
			log.Fatal(err)
		}
	}
	db.SingularTable(true)
	db.LogMode(true)
	initTables()
	initMyCallbacks()
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

func initTables() {
	checkTableData(&model.Admin{})
	checkTableData(&model.AdminOnline{})
	checkTableData(&model.Auth{})
	checkTableData(&model.LoginInfo{})
	checkTableData(&model.OperLog{})
	checkTableData(&model.Role{})
	checkTableData(&model.RoleAuth{})
	checkTableData(&model.SysConf{})

	//if len(ModelGroup) > 0 {
	//	for _, m := range ModelGroup {
	//		checkTableData(m)
	//	}
	//}
}

func checkTableData(tb interface{}) {
	if db.HasTable(tb) == false {
		db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(tb)
		var sqlName string
		if _, ok := tb.(*model.Admin); ok {
			sqlName = "create-admin"
		} else if _, ok := tb.(*model.Auth); ok {
			sqlName = "create-auth"
		} else if _, ok := tb.(*model.Role); ok {
			sqlName = "create-role"
		} else if _, ok := tb.(*model.RoleAuth); ok {
			sqlName = "create-role-auth"
		}
		if sqlName != "" {
			initData(sqlName)
		}
	}
}

func initData(sqlName string) {
	dot, err := dotsql.LoadFromFile("database/base_system.sql")
	if err != nil {
		global.ZapLog.Fatal("无法加载初始数据，请检查data文件夹下是否存在数据信息")
	}
	_, err = dot.Exec(db.DB(), sqlName)
	if err != nil {
		global.ZapLog.Warn("执行" + sqlName + "失败，" + err.Error())
		return
	}
}

func initMyCallbacks() {
	db.Callback().Create().Replace("gorm:update_time_stamp", model.ForBeforeCreate)
	db.Callback().Update().Replace("gorm:update_time_stamp", model.ForBeforeUpdate)
}
