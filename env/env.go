package env

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jiankunking/novel-crawler/pkg/util"
)

var (
	// 剑来起始章节
	JIAN_LAI_START_SECTION = 0

	// dingtalk hook url 地址
	DING_TALK_HOOK = ""

	// 茅山遗孤起始章节
	MAO_SHAN_YI_GU_START_SECTION = 0

	// 牧神记起始章节
	//MU_SHEN_JI_START_SECTION = 0

	GRAB_INTERVAL = "0 0/5 * * * ?"
)

func init() {
	var err error
	index := os.Getenv("JIAN_LAI_START_SECTION")
	if !util.IsEmpty(index) {
		JIAN_LAI_START_SECTION, err = strconv.Atoi(index)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("jianlai:" + index)

	index = os.Getenv("MAO_SHAN_YI_GU_START_SECTION")
	if !util.IsEmpty(index) {
		MAO_SHAN_YI_GU_START_SECTION, err = strconv.Atoi(index)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("maoshanyigu:" + index)

	//index = os.Getenv("MU_SHEN_JI_START_SECTION")
	//if !util.IsEmpty(index) {
	//	MU_SHEN_JI_START_SECTION, err = strconv.Atoi(index)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}
	//fmt.Println("mushenji:" + index)

	accessToken := os.Getenv("ACCESS_TOKEN")
	if !util.IsEmpty(accessToken) {
		DING_TALK_HOOK = "https://oapi.dingtalk.com/robot/send?access_token=" + accessToken
	}
	fmt.Println(DING_TALK_HOOK)

	grabInterval := os.Getenv("GRAB_INTERVAL")
	if !util.IsEmpty(grabInterval) {
		GRAB_INTERVAL = grabInterval
	}
	fmt.Println("定时间隔：" + GRAB_INTERVAL)

}
