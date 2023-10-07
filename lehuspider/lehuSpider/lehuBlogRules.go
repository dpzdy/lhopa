package lehuSpider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	cregex "github.com/mingrammer/commonregex"
	"github.com/tebeka/selenium"
	"log"
	"net/http"
	"os"
	"regexp"
	"spider/items"
	"strings"
	"time"
)

type LehuIndexItem items.LehuIndexItem

const (
	chromeDriverPath = "D:\\GoProject\\WeiboSpiderGo\\browser\\chromedriver.exe"
	port             = 8080
)

// 获取评论 iframe 路径
func getCmturl(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	time.Sleep(time.Second * 2)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	cmtUrl, _ := doc.Find(".box #comment_frame").Attr("src")
	//fmt.Printf("inner function : %s \n", cmtUrl)
	return cmtUrl
}

// 获取评论 返回评论切片
func getCmt(url string, service *selenium.Service, driver selenium.WebDriver) []string {
	cmturl := getCmturl(url)

	targetUrl := "https:" + cmturl
	err := driver.Get(targetUrl)
	if err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}


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

	cmtList := []string{}
	doc.Find("div.bcmtlst ul li.s-bd2").Each(func(i int, s *goquery.Selection) {
		content := s.Find("span.s-fc4").Text()
		content = strings.ReplaceAll(content, " ", "")
		content = strings.ReplaceAll(content, "\n", "")
		cmtList = append(cmtList, content)
	})

	return cmtList
}

// 解析乐乎单个博客帖子 返回对应的单个结构体
func ParseContent(item string, service *selenium.Service, driver selenium.WebDriver) LehuIndexItem {
	lehuItem := LehuIndexItem{}

	resp, err := http.Get(item)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//1.avatar
	Avatar, _ := doc.Find("div.g-sd").Find("img").Attr("src")
	lehuItem.Avatar = Avatar
	//2.author
	Author := doc.Find("div.g-sd").Find("h1.m-ttl").Text()
	lehuItem.Author = strings.TrimSpace(Author)
	doc.Find("div.postinner").Each(func(i int, s *goquery.Selection) {
		//4.content
		Content := strings.TrimSpace(s.Find("div.text").Text())
		lehuItem.Content = Content
		//fmt.Println(content)
		//5.hot
		Hot := s.Find("div.meta .hot").Text()
		regexp, _ := regexp.Compile("[0-9]+")
		Hot = regexp.FindString(Hot)

		lehuItem.Hot = Hot
		//6.cmtNum
		CmtNum := s.Find("div.meta .cmt").Text()
		//str := "Today is Tuesday!"
		CmtNum = regexp.FindString(CmtNum)
		//fmt.Println(regexp.FindString(cmtNum))
		lehuItem.CmtNum = CmtNum
		//7.Date
		Date := s.Find("div.meta").Find("a.date").Text()
		Date = cregex.Date(Date)[0]
		//fmt.Println(cregex.Date(Date)[0])
		lehuItem.Date = Date
		//8.ImgURL
		ImgURL, _ := s.Find("div.pic").Find("img").Attr("src")

		//fmt.Println(ImgURL)
		lehuItem.ImgURL = ImgURL
		var Tags = map[string]string{}
		//9.Tags
		s.Find("div.tags").Find("a.tag").Each(func(i int, selection *goquery.Selection) {
			name := selection.Text()
			href, _ := selection.Attr("href")
			href = strings.ReplaceAll(href, " ", "")
			Tags[name] = href
		})
		lehuItem.Tags = Tags
	})
	cmt := getCmt(item, service, driver)
	lehuItem.Cmt = cmt

	//jsonLehuItem, err := json.Marshal(lehuItem)

	if err != nil {
		fmt.Println("结构体转换Json错误")
		os.Exit(-1)
	}

	//return string(jsonLehuItem)
	return lehuItem
}
