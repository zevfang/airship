package app

import (
	"fmt"
	"github.com/gocolly/colly"
)

// 微博热点爬取
type WeiBo struct {
	Collector colly.Collector
}

func NewCollector() WeiBo {
	return WeiBo{
		Collector: *colly.NewCollector(),
	}
}

func (w *WeiBo) Login() {
	w.Collector.Visit("http://ent.sina.com.cn/")
}

func (w *WeiBo) QueryHot() {
	w.Collector.OnRequest(func(request *colly.Request) {
		fmt.Println("Visiting", request.URL)
	})
	w.Collector.OnResponse(func(response *colly.Response) {
		//fmt.Println(string(response.Body))
	})
	w.Collector.OnHTML("#pl_top_realtimehot > table > tbody ", func(e *colly.HTMLElement) {
		e.ForEach("#pl_top_realtimehot > table > tbody > tr", func(i int, el *colly.HTMLElement) {
			id := el.DOM.Find("td.td-01").Text()
			href, _ := el.DOM.Find("td.td-02 a").Attr("href")
			target := el.DOM.Find("td.td-02 a").Text()
			hot := el.DOM.Find("td.td-03").Text()
			fmt.Println(id, target, href, hot)
		})
	})
	w.Collector.Visit("https://s.weibo.com/top/summary?cate=realtimehot")
}

func (w *WeiBo) Run() {

}
