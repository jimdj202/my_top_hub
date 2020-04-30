package spiders

import (
	"encoding/json"
	"fmt"
	"hub/src/app/db/model"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//TODO
type ResponseDataTianYaRow struct {
	Item string
	Item_name string
	Id string
	Tpye string
	Count string
	Top_count string
	Title string
	Url string
	//Extend string
	Author_id string
	Author_name string
	Reply_time string
	Media string
	Time string
	Pics []string
	Content string
}

type ResponseDataTianYa struct {
	Success string
	Code string
	Message string
	Data struct{
		PageCount string
		Rows []ResponseDataTianYaRow
	}
}

func (s *Spider) GetTianYa() []model.Item{
	typeDomainID := runFuncName()
	fmt.Println("Spider run:", typeDomainID)
	typeDomainID = strings.Split(typeDomainID,"Get")[1]
	var items []model.Item
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	//url := "http://bbs.tianya.cn/list.jsp?item=funinfo&grade=3&order=1"
	url := "https://bbs.tianya.cn/api?method=bbs.ice.getHotArticleList&params.pageSize=50&params.pageNum=1&var=apiData&_r=0.43790093559729804&_=1587869081669"
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}

	request.Header.Add("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36`)
	request.Header.Add("Upgrade-Insecure-Requests", `1`)
	request.Header.Add("Referer", `http://bbs.tianya.cn/list.jsp?item=funinfo&grade=3&order=1`)
	request.Header.Add("Host", `bbs.tianya.cn`)

	res, err := client.Do(request)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	defer res.Body.Close()

	var resData ResponseDataTianYa
	//err = json.NewDecoder(res.Body).Decode(&resData)
	str, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(str))
	str2 := strings.Replace(string(str),"var apiData = ","",1)
	err = json.Unmarshal([]byte(str2), &resData)

	if err == nil{
		for i,v := range resData.Data.Rows{
			//allData = append(allData, map[string]interface{}{"title": text, "url": url})
			b,_ := json.Marshal(v)
			num,_:= strconv.Atoi(v.Top_count)
			imageUrl := ""
			if len (v.Pics) >0 {
				imageUrl = v.Pics[0]
			}
			oneLine := model.Item{
				Index: i,
				Title:      v.Title,
				Url:        v.Url,
				ImageUrl:   imageUrl,
				TypeDomain: "天涯",
				TypeDomainID: typeDomainID,
				TypeFilter: "",
				CommentNum: num ,
				Desc: v.Content,
				Extra: string(b),
				//Date:       time.Time{},
				//CreatedAt:  time.Time{},
				//UpdatedAt:  time.Time{},
				DeletedAt:  nil,
			}
			items = append(items, oneLine)

		}
	}

	//document, err := goquery.NewDocumentFromReader(res.Body)
	//if err != nil {
	//	fmt.Println("抓取" + s.Name + "失败")
	//	return items
	//}
	//document.Find("table tr").Each(func(i int, selection *goquery.Selection) {
	//	s := selection.Find("td a").First()
	//	url, boolUrl := s.Attr("href")
	//	text := s.Text()
	//	comNum := selection.Find(".HotItem-metrics").Text()
	//	reg, _ := regexp.Compile("\\d+")
	//	comNum2 := reg.Find([]byte(comNum))
	//	comNum3, _ := strconv.Atoi(string(comNum2))
	//	imgUrl, _ := selection.Find(".HotItem-img img").Attr("src")
	//	if boolUrl {
	//		//allData = append(allData, map[string]interface{}{"title": text, "url": url})
	//		oneLine := model.Item{
	//			Index: i,
	//			Title:      text,
	//			Url:        url,
	//			ImageUrl:   imgUrl,
	//			TypeDomain: "天涯",
	//			TypeFilter: "",
	//			CommentNum: comNum3 * 10000,
	//			//Date:       time.Time{},
	//			//CreatedAt:  time.Time{},
	//			//UpdatedAt:  time.Time{},
	//			DeletedAt:  nil,
	//		}
	//		items = append(items, oneLine)
	//	}
	//})
	return items
}
