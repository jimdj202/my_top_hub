package spiders

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"hub/src/app/db/model"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)
// 微博
func(s *Spider) GetWeiBo() []model.Item{
	fmt.Println("Spider run:", "WeiBo")
	url := "https://s.weibo.com/top/summary"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	var items []model.Item
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	//request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	defer res.Body.Close()
	//str, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	//b,_:= ioutil.ReadAll(res.Body)
	//fmt.Println(string(b))
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}

	document.Find("tbody tr").Each(func(i int, selection *goquery.Selection) {
		url, boolUrl := selection.Find("a").Attr("href")
		text := selection.Find("a").Text()
		comNum,_:= strconv.Atoi(selection.Find("span").Text())
		textLock := selection.Find("a em").Text()
		text = strings.Replace(text, textLock, "", -1)
		if boolUrl {
			oneLine := model.Item{
				Index: i,
				Title:      text,
				Url:        "https://s.weibo.com" + url,
				ImageUrl:   "",
				TypeDomain: "微博",
				TypeFilter: "",
				CommentNum: comNum,
				//Date:       time.Time{},
				CreatedAt:  time.Time{},
				UpdatedAt:  time.Time{},
				DeletedAt:  nil,
			}
			//allData = append(allData, map[string]interface{}{"title": text, "url": "https://s.weibo.com" + url})
			items = append(items, oneLine)
		}
	})
	return items
}