package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Charger struct {
	Id                      string     `gorm:"type:varchar(36);primary_key" json:"id"`
	Name                    string     `gorm:"type:varchar(32)" json:"name"`
	RatedVoltage            float64    `gorm:"type:decimal(10,2);comment:'unit(v)'" json:"rated_voltage"`
	CreatedAt               *time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt               *time.Time `gorm:"type:timestamp" json:"updated_at"`
	DeletedAt               *time.Time `gorm:"type:timestamp" json:"deleted_at"`
	Province                string     `gorm:"type:varchar(128)" json:"province"`
}

func main() {
	db, err := gorm.Open("mysql", "root:xxxxx@tcp(127.0.0.1:3306)/db_name?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	defer db.Close()
	if err != nil {
		panic("failed to connect database")
	}

	type Result struct {
		Province string
		Total    uint64
	}

	results := make([]Result, 0)
	err = db.Table("charger").Where("deleted_at is null").Select("province, sum(1) as total").Group("province").Scan(&results).Error
	if err != nil {
		panic(err)
	}

	fmt.Println("results ", len(results))
}

// 按照省统计每个省的充电桩总数
// 使用scan与定义的struct从Rows中获取多列数据，如果只是获取一列数据那么使用Pluck即可
