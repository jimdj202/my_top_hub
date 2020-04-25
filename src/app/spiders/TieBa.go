package spiders

import (
	"encoding/json"
	"fmt"
	"hub/src/app/db/model"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type TieBaList struct {
	Topic_id int
	Topic_name string
	Topic_desc string
	Abstract string
	Topic_pic string
	Tag int
	Discuss_num int
	Idx_num int
	Create_time int
	Content_num int
	Topic_avatar string
	Topic_url string
	Topic_default_avatar string
}

type ResponseDataTieBa struct {
	Data struct{
		Bang_topic struct{
			Topic_list []TieBaList
		}
	}
}

func (s *Spider) GetTieBa() []model.Item{
	fmt.Println("Spider run:", "TieBa")
	var items []model.Item
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	url := "http://tieba.baidu.com/hottopic/browse/topicList"
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

	var resData ResponseDataTieBa
	str, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	err = json.Unmarshal(str, &resData)
	if err == nil {
		for i,v := range resData.Data.Bang_topic.Topic_list{

			oneLine := model.Item{
				Index: i,
				Title:      v.Topic_name,
				Url:        v.Topic_url,
				ImageUrl:   v.Topic_pic,
				TypeDomain: "百度贴吧",
				TypeFilter: "",
				CommentNum: v.Discuss_num,
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
	//document, err := goquery.NewDocumentFromReader(res.Body)
	//if err != nil {
	//	fmt.Println("抓取" + s.Name + "失败")
	//	return items
	//}
	//document.Find(".HotItem").Each(func(i int, selection *goquery.Selection) {
	//	url, boolUrl := selection.Find(".HotItem-content a").Attr("href")
	//	text := selection.Find("h2").Text()
	//	comNum := selection.Find(".HotItem-metrics").Text()
	//	reg, _ := regexp.Compile("\\d+")
	//	comNum2 := reg.Find([]byte(comNum))
	//	comNum3, _ := strconv.Atoi(string(comNum2))
	//	imgUrl, _ := selection.Find(".HotItem-img img").Attr("src")
	//	if boolUrl {
	//		//allData = append(allData, map[string]interface{}{"title": text, "url": url})
	//		oneLine := model.Item{
	//			Title:      text,
	//			Url:        url,
	//			ImageUrl:   imgUrl,
	//			TypeDomain: "百度贴吧",
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
