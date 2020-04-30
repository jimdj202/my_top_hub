package spiders

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"hub/src/app/db/model"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (s *Spider) Get36Kr() []model.Item{
	typeDomainID := runFuncName()
	fmt.Println("Spider run:", typeDomainID)
	typeDomainID = strings.Split(typeDomainID,"Get")[1]
	var items []model.Item
	timeout := time.Duration(20 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	url := "https://36kr.com/hot-list/catalog"
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}

	request.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	request.Header.Add("Host", `36kr.com`)
	request.Header.Add("Referer", `https://36kr.com/`)

	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	defer res.Body.Close()
	//str, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	document.Find(".article-wrapper").Each(func(i int, selection *goquery.Selection) {
		text := selection.Find("a.article-item-title").Text()
		url, boolUrl := selection.Find(".article-item-description").Attr("href")
		descText := selection.Find(".article-item-description").Text()
		//imgUrl, _ := selection.Find("img").Attr("src")
		comNum := selection.Find("span.kr-flow-bar-hot span").Text()

		reg, _ := regexp.Compile("\\d+")
		comNum2 := reg.Find([]byte(comNum))
		comNum3, _ := strconv.Atoi(string(comNum2))
		//imgUrl, _ := selection.Find(".HotItem-img img").Attr("src")
		//extra := selection.Find("div.f6 span span").Text()
		if boolUrl {
			//allData = append(allData, map[string]interface{}{"title": text, "url": url})
			oneLine := model.Item{
				Index: i,
				Title:      text,
				Url:        url,
				//ImageUrl:   imgUrl,
				TypeDomain: "36Kr",
				TypeDomainID: typeDomainID,
				TypeFilter: "",
				CommentNum: comNum3 ,
				Desc: descText,
				//Extra: extra,
				//Date:       time.Time{},

				DeletedAt:  nil,
			}
			items = append(items, oneLine)
		}
	})
	return items
}
