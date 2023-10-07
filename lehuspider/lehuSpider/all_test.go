package lehuSpider

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"spider/db"
	"strings"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {

	//启动 chrome 及简单配置

	var opts []selenium.ServiceOption
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}

	// 禁止加载图片，加快渲染速度
	imgCaps := map[string]interface{}{
		//"profile.managed_default_content_settings.images": 2,
	}

	chromeCaps := chrome.Capabilities{
		Prefs: imgCaps,
		Path:  "",
		Args: []string{
			//"--headless",
			"--start-maximized",
			//"--window-size=1200x600",
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36",
			"--disable-gpu",
			"--disable-impl-side-painting",
			"--disable-gpu-sandbox",
			"--disable-accelerated-2d-canvas",
			"--disable-accelerated-jpeg-decoding",
			"--test-type=ui",
		},
	}
	caps.AddChrome(chromeCaps)

	// 启动 chromedriver server
	service, err := selenium.NewChromeDriverService("D:\\GoProject\\WeiboSpiderGo\\browser\\chromedriver.exe", port, opts...)
	if err != nil {
		log.Printf("Error starting the ChromeDriver server: %v", err)
		return
	}
	defer service.Stop()

	//打开一个网页
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Println(err)
		return
	}
	defer wd.Quit()
	//然后加载URL
	err = wd.Get("https://www.lofter.com/front/login/")
	if err != nil {
		log.Println(fmt.Sprintf("Failed to load page: %s\n", err))
	}

	// 判断是否加载完成
	jsRt, err := wd.ExecuteScript("return document.readyState", nil)
	if err != nil {
		log.Println("exe js err", err)
	}
	fmt.Println("jsRt", jsRt)
	if jsRt != "complete" {
		log.Println("网页加载未完成")
		return
	}
	time.Sleep(2 * time.Second)
	// 点击密码登录
	elems, err := wd.FindElement(selenium.ByXPATH, "//*[@id=\"application\"]/div/div[2]/div[2]/div/div/div[1]/div/a[2]")
	fmt.Println("有几个？")
	elems.Click()
	userName, _ := elems.FindElement(selenium.ByXPATH, "//*[@id=\"application\"]/div/div[2]/div[2]/div/div/div[2]/div[2]/form/div[1]/div[1]/input")
	userName.SendKeys("18845649928")
	password, _ := elems.FindElement(selenium.ByXPATH, "//*[@id=\"application\"]/div/div[2]/div[2]/div/div/div[2]/div[2]/form/div[1]/div[2]/input")
	password.SendKeys("Cx18845649928")
	agree, _ := elems.FindElement(selenium.ByXPATH, "//*[@id=\"application\"]/div/div[2]/div[2]/div/div/div[2]/div[2]/form/div[1]/div[3]/span[1]")
	agree.Click()
	login, _ := elems.FindElement(selenium.ByXPATH, "//*[@id=\"application\"]/div/div[2]/div[2]/div/div/div[2]/div[2]/form/div[1]/button")
	login.Click()
	for i := 0; i < 2; i++ {
		fmt.Println("执行了")
		wd.ExecuteScript("window.scrollTo(window.pageXOffset, document.body.scrollHeight)", nil)
		time.Sleep(time.Second * 2)
	}

	frameHtml, err := wd.PageSource()
	if err != nil {
		log.Println(err)
		return
	}

	var doc *goquery.Document
	doc, err = goquery.NewDocumentFromReader(bytes.NewReader([]byte(frameHtml)))
	if err != nil {
		log.Println(err)
		return
	}
	urls := []string{}
	doc.Find("div.mlistcnt").Each(func(divNum int, divItem *goquery.Selection) {
		// do something
		fmt.Printf("第 %d 个 :  ", divNum)
		url, _ := divItem.Find("a.isayc").Attr("href")
		//fmt.Println(name)
		urls = append(urls, url)
	})

	//for _, item := range urls {
	//	//ParseContent(item)
	//}

}

func TestCmt(t *testing.T) {
	url := getCmturl("https://lofguancha.lofter.com/post/749935d7_2b8630cb9")
	url = "https:" + url
	driverPath := "../browser/chromedriver.exe" //准备工作中下载driver
	crawler := Crawler{}
	service, driver := crawler.StartChrome(driverPath)

	targetUrl := url
	err := driver.Get(targetUrl)
	if err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}
	for i := 0; i < 2; i++ {
		fmt.Printf("我滚了第%d次\n", i)
		driver.SetAlertText("我滚了一次")
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
	service.Stop() // 停止chromedriver
	driver.Quit()  // 关闭浏览器

	for i, s := range cmtList {
		fmt.Println(i, s)
	}

}

func TestGeturls(t *testing.T) {
	driverPath := "../browser/chromedriver.exe" //准备工作中下载driver
	crawler := Crawler{}
	service, driver := crawler.StartChrome(driverPath)
	defer service.Stop() // 停止chromedriver
	defer driver.Quit()  // 关闭浏览器
	targetUrl := "https://lofguancha.lofter.com/view"
	err := driver.Get(targetUrl)
	if err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}

	//放在函数中取消注释
	//defer func() {
	//	service.Stop() // 停止chromedriver
	//	driver.Quit()  // 关闭浏览器
	//}()

	for i := 0; i < 2; i++ {
		fmt.Printf("我滚了第%d次\n", i)
		driver.SetAlertText("我滚了一次")
		driver.ExecuteScript("window.scrollTo(window.pageXOffset, document.body.scrollHeight)", nil)
		time.Sleep(time.Duration(2) * time.Second)

	}

	result, err := driver.PageSource()
	// 将结果写入goquery中，以便用css选择器过滤标签
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(result))
	if err != nil {
		panic(fmt.Sprintf("Failed: %s\n", err))
	}

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

	//return lofguanchaList
	//service.Stop() // 停止chromedriver
	//driver.Quit()  // 关闭浏览器
	url := "mongodb://localhost:27017"
	collection := db.ConnectCollection(url, "lehu", "blog1")

	for i, url := range lofguanchaList {
		time.Sleep(time.Duration(5) * time.Second)
		eachJson := ParseContent(url, service, driver)
		db.InsertOne(collection, eachJson)
		fmt.Println(i, eachJson)
	}

}
