package util

import (
	"strings"
)

//func PrintJson(obj interface{}) {
//	jsons, errs := json.Marshal(obj)
//	if errs != nil {
//		fmt.Println(errs.Error())
//	}
//	fmt.Println(string(jsons))
//}

func IsEmpty(str string) bool {
	if str == "" || len(strings.TrimSpace(str)) == 0 {
		return true
	}
	return false
}

func TrimString(item string) string {
	item = strings.Replace(item, "亿", "00000000", -1)
	item = strings.Replace(item, "万", "0000", -1)
	item = strings.Replace(item, "仟", "000", -1)
	item = strings.Replace(item, "佰", "00", -1)
	item = strings.Replace(item, "拾", "0", -1)

	item = strings.Replace(item, "千", "000", -1)
	item = strings.Replace(item, "百", "00", -1)
	item = strings.Replace(item, "十", "0", -1)

	item = strings.Replace(item, "零", "0", -1)
	item = strings.Replace(item, "一", "1", -1)
	item = strings.Replace(item, "二", "2", -1)
	item = strings.Replace(item, "三", "3", -1)
	item = strings.Replace(item, "四", "4", -1)
	item = strings.Replace(item, "五", "5", -1)
	item = strings.Replace(item, "六", "6", -1)
	item = strings.Replace(item, "七", "7", -1)
	item = strings.Replace(item, "八", "8", -1)
	item = strings.Replace(item, "九", "9", -1)

	return item
}
