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

func (s *Spider) GetDouBanMovie() []model.Item{
	fmt.Println("Spider run:", runFuncName())
	var items []model.Item
	timeout := time.Duration(20 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	url := "https://movie.douban.com/"
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}

	request.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)

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
	document.Find("a.item").Each(func(i int, selection *goquery.Selection) {

		url, boolUrl := selection.Attr("href")
		text := selection.Find("p").Text()

		//descText := selection.Find("p").Text()
		comNum ,_:= selection.Find("span.icon-message").Attr("data-origincount")
		//comNum = strings.ReplaceAll(comNum,",","")
		reg, _ := regexp.Compile("\\d+")
		comNum2 := reg.Find([]byte(comNum))
		comNum3, _ := strconv.Atoi(string(comNum2))
		imgUrl, _ := selection.Find("img.lazyload").Attr("data-src")
		//extra := selection.Find("div.f6 span span").Text()
		if boolUrl {
			//allData = append(allData, map[string]interface{}{"title": text, "url": url})
			oneLine := model.Item{
				Index: i,
				Title:      text,
				Url:        "https://www.qdaily.com/" + url,
				ImageUrl:   imgUrl,
				TypeDomain: "QDaily",
				TypeFilter: "",
				CommentNum: comNum3 ,
				//Desc: descText,
				//Extra: extra,
				//Date:       time.Time{},

				DeletedAt:  nil,
			}
			items = append(items, oneLine)
		}
	})
	return items
}
