package spiders

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"hub/src/app/db/model"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

//https://data.wxb.com/rankArticle

func (s *Spider) GetWeiXin() []model.Item{
	fmt.Println("Spider run:", "WeiXin")
	var items []model.Item
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	url := "https://data.wxb.com/rankArticle"
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36`)

	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	defer res.Body.Close()
	str, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(str))
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	document.Find(".ant-table-row-level-0").Each(func(i int, selection *goquery.Selection) {
		url, boolUrl := selection.Find("a.title-text").Attr("href")
		text := selection.Find("a.title-text").Text()
		author := "[" +  selection.Find(".rank-table-gzh").Text() + "]"
		c := selection.Children().Get(3).Data
		fmt.Println(c)
		comNum := selection.Find(".HotItem-metrics").Text()
		reg, _ := regexp.Compile("\\d+")
		comNum2 := reg.Find([]byte(comNum))
		comNum3, _ := strconv.Atoi(string(comNum2))
		imgUrl, _ := selection.Find(".HotItem-img img").Attr("src")
		if boolUrl {
			//allData = append(allData, map[string]interface{}{"title": text, "url": url})
			oneLine := model.Item{
				Title:      author + text,
				Url:        url,
				ImageUrl:   imgUrl,
				TypeDomain: "微信",
				TypeFilter: "",
				CommentNum: comNum3 * 10000,
				//Date:       time.Time{},
				CreatedAt:  time.Time{},
				UpdatedAt:  time.Time{},
				DeletedAt:  nil,
			}
			items = append(items, oneLine)
		}
	})
	return items
}
