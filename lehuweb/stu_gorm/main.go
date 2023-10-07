package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Article struct {
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

type Info struct {
	Id        int    `stu_gorm:"type:int;not null" json:"id"`
	CityName  string `stu_gorm:"type:varchar(255)" json:"cityName"`
	Title     string `stu_gorm:"type:varchar(255)" json:"title"`
	Price     string `stu_gorm:"type:decimal(10,2)" json:"price"`
	Kilometer string `stu_gorm:"type:decimal(10,2)" json:"kilometer"`
	Year      string `stu_gorm:"type:int" json:"year"`
}

func main() {
	in()
}

func ar() {
	DbUser := "root"
	DbPassWord := "11111"
	DbHost := "127.0.0.1"
	DbPort := "3306"
	DbName := "ginblog"
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DbUser,
		DbPassWord,
		DbHost,
		DbPort,
		DbName,
	)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err == nil {
		fmt.Println("数据库链接成功")
	}
	db.AutoMigrate(&Article{})
	var article Article
	db.Raw("select title,img from article").Find(&article)
	fmt.Println(article)
	db.Where("cid=?", 1).Find(&article)
	fmt.Println(article)
}

func in() {
	DbUser := "root"
	DbPassWord := "11111"
	DbHost := "127.0.0.1"
	DbPort := "3306"
	DbName := "spiders"
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DbUser,
		DbPassWord,
		DbHost,
		DbPort,
		DbName,
	)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err == nil {
		fmt.Println("数据库链接成功")
	}
	//db.AutoMigrate(&Info{})
	var info Info
	var infos []Info

	res := db.Last(&info)
	fmt.Println(res)
	fmt.Println(info)
	db.Delete(&info)
	db.Select("id", "city_name").Where("id>?", 25).Find(&infos)
	fmt.Println(infos)
}
