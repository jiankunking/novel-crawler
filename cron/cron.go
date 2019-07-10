package cron

import (
	"fmt"
	"time"

	"github.com/robfig/cron"

	"github.com/jiankunking/novel-crawler/pkg/service"
)

func Minutes() {
	c := cron.New()
	//每5分钟 一次
	c.AddFunc("0 0/5 * * * ?", grab)
	c.Start()
}

func grab() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	go service.CheckJianLaiUpdate()

	go service.CheckMaoShanYiGuUpdate()

	go service.CheckMuShenJiUpdate()

}
