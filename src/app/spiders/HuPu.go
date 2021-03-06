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

func(s *Spider) GetHuPu() []model.Item{
	typeDomainID := runFuncName()
	fmt.Println("Spider run:", typeDomainID)
	typeDomainID = strings.Split(typeDomainID,"Get")[1]
	var items []model.Item
	url := "https://bbs.hupu.com/all-gambia"
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	request.Header.Add("User-Agent", `Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Mobile Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	request.Header.Add("Referer", `https://bbs.hupu.com/`)
	request.Header.Add("Host", `bbs.hupu.com`)
	res, err := client.Do(request)

	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	defer res.Body.Close()
	//str,_ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))

	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	document.Find(".bbsHotPit li").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find(".textSpan a")
		url, boolUrl := s.Attr("href")
		common := selection.Find(("em")).Text()
		reg:= regexp.MustCompile("(\\d*)回复")
		//comNum3 := reg.Find([]byte(common))
		comNum2 := reg.FindStringSubmatch(common)
		comNum3 := 0
		if len(comNum2) > 1 {
			comNum3, _ = strconv.Atoi(comNum2[1])
		}

		text := s.Text()
		if boolUrl {
			oneLine := model.Item{
				Index: i,
				Title:      text,
				Url:        "https://bbs.hupu.com/" + url,
				ImageUrl:   "",
				TypeDomain: "虎扑",
				TypeDomainID: typeDomainID,
				TypeFilter: "",
				CommentNum: comNum3,
				//Date:       time.Time{},
				CreatedAt:  time.Time{},
				UpdatedAt:  time.Time{},
				DeletedAt:  nil,
			}
			//items = append(items, map[string]interface{}{"title": text, "url": "https://bbs.hupu.com/" + url})
			items = append(items, oneLine)
		}
	})
	return items
}