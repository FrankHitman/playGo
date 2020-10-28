package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := gorm.Open("mysql", "root:xxxxx@tcp(127.0.0.1:3306)/db_name?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	defer db.Close()
	if err != nil {
		panic("failed to connect database")
	}

	e := ElectricVehicle{Base: Base{Id: "1"}}
	// default deleted_at is null and will check id = xxxx if id is not default value
	d := db.First(&e)
	fmt.Println(e, d.Error)

}

// when id = "" db.First() will get the first item where deleted_at is null and order by id asc
// when id != "" db.First() will get the first item where deleted_at is null and id = xxx order by id asc
// gorm中有一些默认行为，例如在First()函数中，只要传入的struct实例化中的主键不为空，就会根据主键来查询数据库。
// 还有一些默认行为：
// 查询时候如果未指定排序就会根据id生序排序
// 如果有deleted_at就是软删除，通过回调函数完成。调用First()时候会默认加上"deleted_at is null"的查询条件
// 如果有updated_at就会在每次更新时候自动回调更新updated_at日期


// 回调函数调用
// func (scope *Scope) callCallbacks(funcs []*func(s *Scope)) *Scope {
// 	defer func() {
// 		if err := recover(); err != nil {
// 			if db, ok := scope.db.db.(sqlTx); ok {
// 				db.Rollback()
// 			}
// 			panic(err)
// 		}
// 	}()
// 	for _, f := range funcs {
// 		(*f)(scope)     			// 执行回调函数
// 		if scope.skipLeft {
// 			break
// 		}
// 	}
// 	return scope
// }