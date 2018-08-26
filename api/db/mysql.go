package db

import (
	"github.com/go-xorm/xorm"
	"fmt"
	"github.com/go-xorm/core"
	_ "github.com/go-sql-driver/mysql"
)

var engine *xorm.Engine

func init() {
	var err error

	engine, err = xorm.NewEngine("mysql", "root:admin@/tools?charset=utf8")

	if err != nil {
		fmt.Println(err)
	}
	if err := engine.Ping(); err != nil{
		fmt.Println(err)
	}

	engine.SetMaxIdleConns(5)
	engine.SetMaxOpenConns(5)

	engine.SetTableMapper(core.SnakeMapper{})
}

func GetMysqlLink() *xorm.Engine {
	return engine;
}
