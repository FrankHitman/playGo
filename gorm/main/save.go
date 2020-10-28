package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Base struct {
	Id        string     `gorm:"type:varchar(36);primary_key" json:"id"`
	CreatedAt *time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:timestamp" json:"updated_at"`
	// DeletedAt *time.Time `gorm:"type:timestamp"`
}

type ElectricVehicle struct {
	Base
	MobileUserId    string     `gorm:"type:varchar(36);not null" json:"mobile_user_id"`
	CarPlate        string     `gorm:"type:varchar(32)" json:"car_plate"`
	Brand           string     `gorm:"type:varchar(32)" json:"brand"`
	Model           string     `gorm:"type:varchar(32)" json:"model"`
	DeletedAt       *time.Time `gorm:"type:timestamp" json:"deleted_at"`
}

func main() {
	db, err := gorm.Open("mysql", "root:xxxxx@tcp(127.0.0.1:3306)/db_name?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	defer db.Close()
	if err != nil {
		panic("failed to connect database")
	}

	e := ElectricVehicle{Base: Base{Id: "1"},
		MobileUserId: "d84d9ed5-6f34-432b-9865-c8a6771cd867"}
	d := db.Save(&e)
	fmt.Println(e, d.Error)
}

// Save update value in database, if the value doesn't have primary key, will insert it
// Save will include all fields when perform the Updating SQL, even it is not changed
// when used in create e this will go pass update -> select -> insert 3 progress. =_=
// 如果用save做创建操作，那么性能是比较浪费的，通过debug代码可以知道经历了三个过程：
// 先是通过主键更新，如果更新的行数为0（RowsAffected==0），那么再去执行FirstOrCreate()，
// 尝试根据主键select一把，如果数据库中查询不到（RecordNotFound），最后才是往数据库中写入数据，浪费了两次sql操作。

// reference http://jinzhu.me/gorm/crud.html#update

// 源代码：
// // Save update value in database, if the value doesn't have primary key, will insert it
// func (s *DB) Save(value interface{}) *DB {
//	scope := s.NewScope(value)
//	if !scope.PrimaryKeyZero() {
//		newDB := scope.callCallbacks(s.parent.callbacks.updates).db
//		if newDB.Error == nil && newDB.RowsAffected == 0 {
//			return s.New().FirstOrCreate(value)
//		}
//		return newDB
//	}
//	return scope.callCallbacks(s.parent.callbacks.creates).db
// }

// // FirstOrCreate find first matched record or create a new one with given conditions (only works with struct, map conditions)
// // https://jinzhu.github.io/gorm/crud.html#firstorcreate
// func (s *DB) FirstOrCreate(out interface{}, where ...interface{}) *DB {
//	c := s.clone()
//	if result := s.First(out, where...); result.Error != nil {
//		if !result.RecordNotFound() {
//			return result
//		}
//		return c.NewScope(out).inlineCondition(where...).initialize().callCallbacks(c.parent.callbacks.creates).db
//	} else if len(c.search.assignAttrs) > 0 {
//		return c.NewScope(out).InstanceSet("gorm:update_interface", c.search.assignAttrs).callCallbacks(c.parent.callbacks.updates).db
//	}
//	return c
// }