package cron

import (
	"github.com/robfig/cron"

	"github.com/jiankunking/novel-crawler/env"
	"github.com/jiankunking/novel-crawler/pkg/service"
)

func Minutes() {
	c := cron.New()
	//每5分钟 一次
	c.AddFunc(env.GRAB_INTERVAL, grab)
	c.Start()
}

func grab() {
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " 开始检查是否有更新")

	go service.CheckJianLaiUpdate()

	go service.CheckMaoShanYiGuUpdate()

	//go service.CheckMuShenJiUpdate()

	//fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " 更新检测结束")

}
