package data

import (
	"fmt"
	"strconv"
)

type Info struct {
	Id        int    `stu_gorm:"type:int;not null" json:"id"`
	CityName  string `stu_gorm:"type:varchar(255)" json:"cityName"`
	Title     string `stu_gorm:"type:varchar(255)" json:"title"`
	Price     string `stu_gorm:"type:decimal(10,2)" json:"price"`
	Kilometer string `stu_gorm:"type:decimal(10,2)" json:"kilometer"`
	Year      string `stu_gorm:"type:int" json:"year"`
}

func GetAllInfo() ([]Info, int64) {
	var infos []Info
	var total int64

	db.Raw("select * from info").Find(&infos)
	db.Model(&infos).Count(&total)
	return infos[:20], total

}
func GetInfoByCity(city string) []Info {
	var infos []Info

	db.Where("city_name=?", city).Find(&infos)
	return infos
}
func GetInfosById(left int) []Info {
	var infos []Info

	fmt.Println(left)
	db.Where("id > ?", strconv.Itoa(left)).Find(&infos)
	fmt.Println(infos)
	return infos
}

//db.Raw("delete from sp_qc_car where id > 50").Find(&info)
//db.Raw("select id,city_name from info").Find(&infos)
//db.Raw("insert  into sp_qc_car (id , city_name) values (51,'北京') ").Find(&info)
//db.Raw("update sp_qc_car set city_name = '山东' where id = 51").Find(&info)
