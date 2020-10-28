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

	var traceabilityNums []string
	d := db.Table("charger").Joins("inner join private_charger "+
		"on private_charger.charger_id = charger.id "+
		"and private_charger.mobile_user_id = ? "+
		"and private_charger.deleted_at is null "+
		"and charger.deleted_at is null", "xxxxxx").Pluck("fk_traceability_num", &traceabilityNums)
	fmt.Println(traceabilityNums, d.Error)

}


// pluck中的value值必须是数组指针，否则会报错，例外这个数组可以是初始化过的（make([]string,0)）也可以是只声明未初始化过的（var xxxx []string），
// 参考reflect.go实验
// func (scope *Scope) pluck(column string, value interface{}) *Scope {
// 	dest := reflect.Indirect(reflect.ValueOf(value))
// 	if dest.Kind() != reflect.Slice {
// 		scope.Err(fmt.Errorf("results should be a slice, not %s", dest.Kind()))
// 		return scope
// 	}
//
// 	if dest.Len() > 0 {
// 		dest.Set(reflect.Zero(dest.Type()))
// 	}
//
// 	if query, ok := scope.Search.selects["query"]; !ok || !scope.isQueryForColumn(query, column) {
// 		scope.Search.Select(column)
// 	}
//
// 	rows, err := scope.rows()
// 	if scope.Err(err) == nil {
// 		defer rows.Close()
// 		for rows.Next() {
// 			elem := reflect.New(dest.Type().Elem()).Interface()
// 			scope.Err(rows.Scan(elem))
// 			dest.Set(reflect.Append(dest, reflect.ValueOf(elem).Elem()))
// 		}
//
// 		if err := rows.Err(); err != nil {
// 			scope.Err(err)
// 		}
// 	}
// 	return scope
// }