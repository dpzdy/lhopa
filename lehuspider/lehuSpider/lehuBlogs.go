package lehuSpider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"spider/db"
	"strings"
	"time"
)

// 后期增加diverpath 和 tagetUrl
func BlogSpider() {
	//相对路径使用的是执行路径
	driverPath := "./browser/chromedriver.exe" //准备工作中下载driver
	crawler := Crawler{}
	service, driver := crawler.StartChrome(driverPath)
	defer service.Stop() // 停止chromedriver
	defer driver.Quit()  // 关闭浏览器

	//开始
	targetUrl := "https://lofguancha.lofter.com/view"
	err := driver.Get(targetUrl)
	if err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}

	//复用，提到util
	for i := 0; i < 2; i++ {
		driver.ExecuteScript("window.scrollTo(window.pageXOffset, document.body.scrollHeight)", nil)
		time.Sleep(time.Duration(2) * time.Second)
	}

	result, err := driver.PageSource()
	// 将结果写入goquery中，以便用css选择器过滤标签
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(result))
	if err != nil {
		panic(fmt.Sprintf("Failed: %s\n", err))
	}
	//后期改成常量，LEHU_URL
	rootPath := "https://lofguancha.lofter.com"
	lofguanchaList := []string{}
	fmt.Println(doc.Find(".m-filecnt-1 .list li").Length())
	doc.Find(".m-filecnt-1 .list li").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Find("a").Attr("href")
		url = rootPath + url
		//fmt.Println(url)
		lofguanchaList = append(lofguanchaList, url)
	})
	for index, url := range lofguanchaList {
		fmt.Println(index, url)
	}

	url := "mongodb://localhost:27017"
	collection := db.ConnectCollection(url, "lehu", "blog1")

	for i, url := range lofguanchaList {
		time.Sleep(time.Duration(5) * time.Second)
		eachJson := ParseContent(url, service, driver)
		db.InsertOne(collection, eachJson)
		fmt.Println(i, eachJson)
	}

}
