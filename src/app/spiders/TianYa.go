package spiders

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"hub/src/app/db/model"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

//TODO

func (s *Spider) GetTianYa() []model.Item{
	fmt.Println("Spider run:", "TianYa")
	var items []model.Item
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	url := "http://bbs.tianya.cn/list.jsp?item=funinfo&grade=3&order=1"
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}

	request.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	request.Header.Add("Referer", `http://bbs.tianya.cn/list.jsp?item=funinfo&grade=3&order=1`)
	request.Header.Add("Host", `bbs.tianya.cn`)

	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	document.Find("table tr").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("td a").First()
		url, boolUrl := s.Attr("href")
		text := s.Text()
		comNum := selection.Find(".HotItem-metrics").Text()
		reg, _ := regexp.Compile("\\d+")
		comNum2 := reg.Find([]byte(comNum))
		comNum3, _ := strconv.Atoi(string(comNum2))
		imgUrl, _ := selection.Find(".HotItem-img img").Attr("src")
		if boolUrl {
			//allData = append(allData, map[string]interface{}{"title": text, "url": url})
			oneLine := model.Item{
				Index: i,
				Title:      text,
				Url:        url,
				ImageUrl:   imgUrl,
				TypeDomain: "天涯",
				TypeFilter: "",
				CommentNum: comNum3 * 10000,
				//Date:       time.Time{},
				//CreatedAt:  time.Time{},
				//UpdatedAt:  time.Time{},
				DeletedAt:  nil,
			}
			items = append(items, oneLine)
		}
	})
	return items
}
