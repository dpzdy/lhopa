package data

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
)

type LehuIndexItem struct {
	Avatar  string            `json:"avatar"`
	Author  string            `json:"author"`
	Title   string            `json:"title,omitempty"`
	Content string            `json:"content"`
	Hot     string            `json:"hot"`
	CmtNum  string            `json:"cmtNum"`
	Date    string            `json:"date"`
	ImgURL  string            `json:"imgURL"`
	Tags    map[string]string `json:"Tags"`
	Cmt     []string          `json:"cmt"`
}

func GetAllInfo() []LehuIndexItem {
	fmt.Println("Find:")
	opt := options.Find()
	opt.SetLimit(5)

	col := mgo.Database("lehu").Collection("blog")

	//Find查询多条记录，将会返回一个迭代器，使用完记得关闭
	cor, err := col.Find(context.TODO(), bson.D{{"author", "老福特观察局"}}, opt)
	if err != nil {
		log.Fatalln(err)
	}
	defer cor.Close(context.TODO())

	res := []LehuIndexItem{}
	for cor.Next(context.TODO()) {
		var p LehuIndexItem
		cor.Decode(&p)
		//fmt.Println(p)
		res = append(res, p)
	}
	return res

}

//db.Raw("delete from sp_qc_car where id > 50").Find(&info)
//db.Raw("select id,city_name from info").Find(&infos)
//db.Raw("insert  into sp_qc_car (id , city_name) values (51,'北京') ").Find(&info)
//db.Raw("update sp_qc_car set city_name = '山东' where id = 51").Find(&info)
