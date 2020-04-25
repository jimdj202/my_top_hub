package spiders

import (
	"encoding/json"
	"fmt"
	"hub/src/app/db/model"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)
type ResponseItem struct {
	Id string
	//WxOriginId string
	Wx_origin_id string
	Url string
	Title string
	Index_scores float32
	Read_num interface{}
	Like_num int
	Update_time string
	Content_type int
	Is_original int
	Account string
	Avatar string

}

type ResponseData struct {
	Errcode uint8
	Data []ResponseItem

}

//https://data.wxb.com/rankArticle

func (s *Spider) GetWeiXin() []model.Item{
	fmt.Println("Spider run:", "WeiXin")
	var items []model.Item
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	//https://data.wxb.com/rank/article?baidu_cat=%E6%80%BB%E6%A6%9C&baidu_tag=&page=1&pageSize=50&type=2&order=
	//url := "https://data.wxb.com/rankArticle"
	url := "https://data.wxb.com/rank/article?baidu_cat=%E6%80%BB%E6%A6%9C&baidu_tag=&page=1&pageSize=50&type=2&order="

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

	var resData ResponseData
	str, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	err = json.Unmarshal(str, &resData)

	if err == nil {
		for _,v := range resData.Data{
			reg, _ := regexp.Compile("\\d+")
			var comNum3 int
			switch v.Read_num.(type) {
			case string:
				comNum2 := reg.Find([]byte(v.Read_num.(string)))
				comNum3i, _ := strconv.Atoi(string(comNum2))
				comNum3 = comNum3i * 10000
			case float32:
				comNum3 = int(v.Read_num.(float32))
			}

			oneLine := model.Item{
				Title:      "["+v.Account+"]" + v.Title,
				Url:        v.Url,
				ImageUrl:   v.Avatar,
				TypeDomain: "微信",
				TypeFilter: "",
				CommentNum: comNum3,
				//Date:       time.Time{},
				//CreatedAt:  time.Time{},
				//UpdatedAt:  time.Time{},
				DeletedAt:  nil,
			}
			items = append(items, oneLine)

		}
	}else {
		fmt.Println(err)
	}

	//fmt.Println(string(str))
	//document, err := goquery.NewDocumentFromReader(res.Body)
	//if err != nil {
	//	fmt.Println("抓取" + s.Name + "失败")
	//	return items
	//}
	//document.Find(".ant-table-row-level-0").Each(func(i int, selection *goquery.Selection) {
	//	url, boolUrl := selection.Find("a.title-text").Attr("href")
	//	text := selection.Find("a.title-text").Text()
	//	author := "[" +  selection.Find(".rank-table-gzh").Text() + "]"
	//	c := selection.Children().Get(3).Data
	//	fmt.Println(c)
	//	comNum := selection.Find(".HotItem-metrics").Text()
	//	reg, _ := regexp.Compile("\\d+")
	//	comNum2 := reg.Find([]byte(comNum))
	//	comNum3, _ := strconv.Atoi(string(comNum2))
	//	imgUrl, _ := selection.Find(".HotItem-img img").Attr("src")
	//	if boolUrl {
	//		//allData = append(allData, map[string]interface{}{"title": text, "url": url})
	//		oneLine := model.Item{
	//			Title:      author + text,
	//			Url:        url,
	//			ImageUrl:   imgUrl,
	//			TypeDomain: "微信",
	//			TypeFilter: "",
	//			CommentNum: comNum3 * 10000,
	//			//Date:       time.Time{},
	//			CreatedAt:  time.Time{},
	//			UpdatedAt:  time.Time{},
	//			DeletedAt:  nil,
	//		}
	//		items = append(items, oneLine)
	//	}
	//})
	return items
}
