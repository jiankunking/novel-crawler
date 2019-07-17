package service

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/jiankunking/novel-crawler/env"
	"github.com/jiankunking/novel-crawler/pkg/util"
)

func CheckMuShenJiUpdate() {
	index := "https://www.biquge18.com/book/201/"
	doc, err := goquery.NewDocument(index)
	if err != nil {
		fmt.Println(err)
		return
	}

	title := doc.Find("div.box_con").Find("dd").Eq(0).Text()
	// fmt.Println(title)
	if util.IsEmpty(title) {
		fmt.Println("mushenji title is empty")
		return
	}
	items := strings.Split(title, " ")
	item := items[0]
	// fmt.Println(item)
	startIndex := len("第")
	item = item[startIndex : len(item)-startIndex]
	// fmt.Println(item)
	current := util.TakeChineseNumberFromString(item)

	href, _ := doc.Find("div.box_con").Find("dd").Eq(0).Find("a").Attr("href")
	if current > env.MU_SHEN_JI_START_SECTION {
		env.MU_SHEN_JI_START_SECTION = current
		hrefs := strings.Split(href, "/")
		href = hrefs[len(hrefs)-1]
		req := NewRequest(env.DING_TALK_HOOK, "牧神记已更新："+title, index+href)
		if _, err := req.Send(); err != nil {
			fmt.Printf("Dingtalk send failed cause %s", err.Error())
		}
	}
}
