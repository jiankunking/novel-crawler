package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/jiankunking/novel-crawler/env"
	"github.com/jiankunking/novel-crawler/pkg/util"
)

func CheckJianLaiUpdate() {
	index := "http://www.shuquge.com/txt/8659/index.html"
	doc, err := NewDocumentWithTimeout(index, time.Duration(5)*time.Second)
	//doc, err := goquery.NewDocument(index)
	if err != nil {
		fmt.Println(err)
		return
	}
	title := doc.Find("div.listmain").Find("dd").Eq(0).Text()
	if util.IsEmpty(title) {
		fmt.Println("jianlai title is empty")
		return
	}
	// fmt.Println(title)
	items := strings.Split(title, " ")
	item := items[0]
	//fmt.Println(item)
	startIndex := len("第")
	item = item[startIndex : len(item)-startIndex]
	//fmt.Println(item)

	current := util.TakeChineseNumberFromString(item)
	href, _ := doc.Find("div.listmain").Find("dd").Eq(0).Find("a").Attr("href")
	//fmt.Println(href)
	if current > env.JIAN_LAI_START_SECTION {
		env.JIAN_LAI_START_SECTION = current
		index := index[0 : len(index)-10]
		req := NewRequest(env.DING_TALK_HOOK, "剑来已更新："+title, index+"/"+href)
		if _, err := req.Send(); err != nil {
			fmt.Printf("Dingtalk send failed cause %s", err.Error())
		}
	}

}
