package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func Init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbprot := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbprot == "" {
		dbprot = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbprot + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(User))
}

func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}
