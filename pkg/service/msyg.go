package service

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/jiankunking/novel-crawler/env"
	"github.com/jiankunking/novel-crawler/pkg/util"
)

func CheckMaoShanYiGuUpdate() {
	index := "http://www.biquge.cc/html/307/307566/"
	doc, err := goquery.NewDocument(index)
	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find("div.box_con").Find("dd").Eq(0).Text()
	// fmt.Println(title)
	if util.IsEmpty(title) {
		fmt.Println("maoshanyigu title is empty")
		return
	}
	items := strings.Split(title, ":")
	item := items[0]
	item = strings.Replace(item, " ", "", -1)
	// fmt.Println(item)
	startIndex := len("第")
	item = item[startIndex : len(item)-startIndex]
	// fmt.Println(item)
	current, err := strconv.Atoi(item)
	if err != nil {
		log.Fatal(err)
		return
	}
	href, _ := doc.Find("div.box_con").Find("dd").Eq(0).Find("a").Attr("href")
	if current > env.MAO_SHAN_YI_GU_START_SECTION {
		env.MAO_SHAN_YI_GU_START_SECTION = current
		req := NewRequest(env.DING_TALK_HOOK, "茅山遗孤已更新："+title, index+href)
		if _, err := req.Send(); err != nil {
			fmt.Printf("Dingtalk send failed cause %s", err.Error())
		}
	}
}
