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

func (s *Spider) GetGitHub() []model.Item{
	typeDomainID := runFuncName()
	fmt.Println("Spider run:", typeDomainID)
	typeDomainID = strings.Split(typeDomainID,"Get")[1]
	var items []model.Item
	timeout := time.Duration(20 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	url := "https://github.com/trending?since=daily"
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

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	document.Find(".Box article").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find(".lh-condensed a")
		//desc := selection.Find(".col-9 .text-gray .my-1 .pr-4")
		//descText := desc.Text()
		url, boolUrl := s.Attr("href")
		text := s.Text()
		text = strings.ReplaceAll(text,"\n","")
		text = strings.ReplaceAll(text," ","")
		descText := selection.Find("p").Text()
		comNum := selection.Find("a.muted-link:nth-child(2)").Text()
		comNum = strings.ReplaceAll(comNum,",","")
		reg, _ := regexp.Compile("\\d+")
		comNum2 := reg.Find([]byte(comNum))
		comNum3, _ := strconv.Atoi(string(comNum2))
		//imgUrl, _ := selection.Find(".HotItem-img img").Attr("src")
		extra := selection.Find("div.f6 span span").Text()
		if boolUrl {
			//allData = append(allData, map[string]interface{}{"title": text, "url": url})
			oneLine := model.Item{
				Index: i,
				Title:      text,
				Url:        "https://github.com" + url,
				//ImageUrl:   imgUrl,
				TypeDomain: "GitHub",
				TypeDomainID: typeDomainID,
				TypeFilter: "",
				CommentNum: comNum3 ,
				Desc: descText,
				Extra: extra,
				//Date:       time.Time{},

				DeletedAt:  nil,
			}
			items = append(items, oneLine)
		}
	})
	return items
}
