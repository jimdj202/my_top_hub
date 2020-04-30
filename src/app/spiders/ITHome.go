package spiders

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"hub/src/app/db/model"
	"io"
	"net/http"
	"strings"
	"time"
)

func (s *Spider) GetITHome() []model.Item{
	typeDomainID := runFuncName()
	fmt.Println("Spider run:", typeDomainID)
	typeDomainID = strings.Split(typeDomainID,"Get")[1]
	var items []model.Item
	timeout := time.Duration(5 * time.Second) //超时时间5s
	client := &http.Client{
		Timeout: timeout,
	}
	url := "https://www.ithome.com/"
	var Body io.Reader
	request, err := http.NewRequest("GET", url, Body)
	if err != nil {
		fmt.Println("抓取" + s.Name + "失败")
		return items
	}
	//request.Header.Add("Cookie", `_zap=41286a80-38b3-4096-bb75-db0ddfdd73dc; d_c0="AKCXnzQZ-hCPTraAao2sQxF9_pK06BheNo4=|1584441730"; _ga=GA1.2.515657607.1584441732; q_c1=646c9d033bf74c628c3967b3b9c0bcdb|1587347613000|1587347613000; __utma=51854390.515657607.1584441732.1587347661.1587347661.1; __utmz=51854390.1587347661.1.1.utmcsr=(direct)|utmccn=(direct)|utmcmd=(none); __utmv=51854390.100--|2=registration_date=20200321=1^3=entry_date=20200321=1; _xsrf=d7173b3f-bbcf-4ff7-93d9-fe1abca324fb; _gid=GA1.2.1202323160.1587720679; Hm_lvt_98beee57fd2ef70ccdd5ca52b9740c49=1587347613,1587347682,1587720678,1587720707; capsion_ticket="2|1:0|10:1587721761|14:capsion_ticket|44:ZTFiNjhiMjUyNjlhNGNlZWEzYjQyNWFjNDc5ZmVlZTg=|340b28b985b1ecfef94f9ef087ab128a5798e12bbcc60d1fa4ccea3b349767f2"; z_c0="2|1:0|10:1587721831|4:z_c0|92:Mi4xZlBrS0ZRQUFBQUFBb0plZk5CbjZFQ2NBQUFDRUFsVk5aa1BLWGdDdE1jSTdPTmRrd2JBV0dEcnk1SVB4WnlUVEpR|40d2feaf6b5d623287b580437aa5eedca475c393de46d4ee4e6aae95ad655893"; _gat_gtag_UA_149949619_1=1; tshl=; tst=h; Hm_lpvt_98beee57fd2ef70ccdd5ca52b9740c49=1587721876; KLBRSID=37f2e85292ebb2c2ef70f1d8e39c2b34|1587721868|1587720669; SESSIONID=0F78sW3Ag8pQbk91JqT8AaisoiytQcCMhvJj8G7kHo6; JOID=UVAQBkiTy9QfctLTS5Xwz6-ELfJcqIDjIQGO7Qr0-KlOFrnhdxlwcEd51N5NnY13VG6xRgUZaxgdWGdErKFgYsM=; osd=U1AdAkuRy9kbcdDTRpHzza-JKfFeqI3nIgOO4A73-qlDErrjdxR0c0V52dpOn416UG2zRggdaBodVWNHrqFtZsA=`)
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
	document.Find(".hot-list .bx ul li").Each(func(i int, selection *goquery.Selection) {
		url, boolUrl := selection.Find("a").Attr("href")
		text := selection.Find("a").Text()
		//comNum := selection.Find(".HotItem-metrics").Text()
		//reg, _ := regexp.Compile("\\d+")
		//comNum2 := reg.Find([]byte(comNum))
		//comNum3, _ := strconv.Atoi(string(comNum2))
		//imgUrl, _ := selection.Find(".HotItem-img img").Attr("src")
		if boolUrl {
			//allData = append(allData, map[string]interface{}{"title": text, "url": url})
			oneLine := model.Item{
				Index: i,
				Title:      text,
				Url:        url,
				//ImageUrl:   imgUrl,
				TypeDomain: "ITHome",
				TypeDomainID: typeDomainID,
				TypeFilter: "",
				//CommentNum: comNum3 * 10000,
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
