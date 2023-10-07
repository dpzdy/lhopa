package lehuSpider

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
)

// 定义一个爬虫结构体
type Crawler struct {
}

// 面向对象思想，将配置定义成方法
func (c Crawler) Config() (opts []selenium.ServiceOption, caps selenium.Capabilities) {
	opts = []selenium.ServiceOption{}
	caps = selenium.Capabilities{
		"browserName": "chrome",
	}
	// 禁止加载图片，加快渲染速度
	imagCaps := map[string]interface{}{
		//"profile.managed_default_content_settings.images": 2,
	}
	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			"--headless", // 设置Chrome无头模式
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬
		},
	}
	caps.AddChrome(chromeCaps)
	return opts, caps
}

// 爬虫启动
func (crawler Crawler) StartChrome(path string) (*selenium.Service, selenium.WebDriver) {
	opts, caps := crawler.Config()
	// 启动chromedriver，端口号可自定义
	service, err := selenium.NewChromeDriverService(path, 9999, opts...)
	if err != nil {
		log.Printf("Error starting the ChromeDriver server: %v", err)
	}

	// 调起chrome浏览器
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9999))
	if err != nil {
		panic(err)
	}
	//defer service.Stop()   // 停止chromedriver
	//defer webDriver.Quit() // 关闭浏览器
	return service, webDriver
}
