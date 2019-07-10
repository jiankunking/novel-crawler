package main

import (
	"github.com/jiankunking/novel-crawler/cron"
)

func main() {
	//启动定时任务
	go cron.Minutes()
	select {}
}
