package spiders

import (
	"fmt"
	"hub/src/app/db/model"
	"io"
	"net/http"
	"time"
)

func(s *Sipder) GetHuPu() *[]model.Item{
	fmt.Println("Spider run:", "HuPu")
	items := &[]model.Item{}
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
	var allData []map[string]interface{}
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("抓取" + spider.DataType + "失败")
		return []map[string]interface{}{}
	}
	document.Find(".bbsHotPit li").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find(".textSpan a")
		url, boolUrl := s.Attr("href")
		text := s.Text()
		if boolUrl {
			allData = append(allData, map[string]interface{}{"title": text, "url": "https://bbs.hupu.com/" + url})
		}
	})
	return allData
}