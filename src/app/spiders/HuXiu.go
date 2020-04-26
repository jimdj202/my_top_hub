package spiders

import (
	"fmt"
	"github.com/json-iterator/go"
	"hub/src/app/db/model"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func (s *Spider) GetHuXiu() []model.Item{
	fmt.Println("Spider run:", runFuncName())
	var items []model.Item
	timeout := time.Duration(20 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	//url := "https://www.huxiu.com/article"
	url := "https://article-api.huxiu.com/web/article/articleList?platform=www"
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}

	request.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	request.Header.Add("Host", `https://www.huxiu.com`)
	request.Header.Add("Referer", `https://www.huxiu.com`)

	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	defer res.Body.Close()

	bodyData, _ := ioutil.ReadAll(res.Body)
	array := jsoniter.Get(bodyData,"data").Get("dataList")
	for i:=0 ; i < array.Size();i++{
		item := array.Get(i)
		url := strings.Replace(item.Get("share_url").ToString(),"https://m","https://www",1)
		oneLine := model.Item{
			Index: i,
			Title:      item.Get("title").ToString(),
			Url:        url,
			ImageUrl:   item.Get("origin_pic_path").ToString(),
			TypeDomain: "HuXiu",
			TypeFilter: "",
			CommentNum: item.Get("count_info").Get("viewnum").ToInt() ,
			Desc: item.Get("summary").ToString(),
			Extra: item.Get("label").ToString(),
			//Date:       time.Time{},
			//CreatedAt:  time.Time{},
			//UpdatedAt:  time.Time{},
			DeletedAt:  nil,
		}
		items = append(items, oneLine)
	}
	//fmt.Println(string(str))
	//document, err := goquery.NewDocumentFromReader(res.Body)
	//if err != nil {
	//	fmt.Println("抓取" + s.Name + "失败")
	//	return items
	//}
	//document.Find("div[data-log=expose]").Each(func(i int, selection *goquery.Selection) {
	//	s := selection.Find("a:nth-of-type(1)")
	//	url, boolUrl := s.Attr("href")
	//	text := s.Find("h5").Text()
	//
	//	descText := selection.Find("p.article-summary").Text()
	//	//comNum := selection.Find("span.article-comments-num").Text()
	//	//comNum = strings.ReplaceAll(comNum,",","")
	//	//reg, _ := regexp.Compile("\\d+")
	//	//comNum2 := reg.Find([]byte(comNum))
	//	//comNum3, _ := strconv.Atoi(string(comNum2))
	//	imgUrl, _ := selection.Find("img").Attr("data-src")
	//	extra := selection.Find("a.article-author").Text()
	//	if boolUrl {
	//		//allData = append(allData, map[string]interface{}{"title": text, "url": url})
	//		oneLine := model.Item{
	//			Index: i,
	//			Title:      text,
	//			Url:        "https://www.huxiu.com" + url,
	//			ImageUrl:   imgUrl,
	//			TypeDomain: "HuXiu",
	//			TypeFilter: "",
	//			//CommentNum: comNum3 ,
	//			Desc: descText,
	//			Extra: extra,
	//			//Date:       time.Time{},
	//
	//			DeletedAt:  nil,
	//		}
	//		items = append(items, oneLine)
	//	}
	//})
	return items
}
