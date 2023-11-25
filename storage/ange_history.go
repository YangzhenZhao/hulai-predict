package storage

import (
	"time"

	"github.com/YangzhenZhao/hulai-predict/dto"
)

func newDate(year int, month time.Month, day int) time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Date(year, month, day, 0, 0, 0, 0, location)
}

var AngeHistoryMap = map[string][]dto.AngeHeros{
	"wei": []dto.AngeHeros{
		{Date: newDate(2023, 2, 3), FirstHero: "张辽", SecondHero: "蔡文姬"},
		{Date: newDate(2023, 3, 3), FirstHero: "司马懿", SecondHero: "曹仁"},
		{Date: newDate(2023, 3, 31), FirstHero: "夏侯惇", SecondHero: "郭嘉"},
		{Date: newDate(2023, 4, 28), FirstHero: "曹丕", SecondHero: "张春华"},
		{Date: newDate(2023, 5, 26), FirstHero: "曹操", SecondHero: "张辽"},
		{Date: newDate(2023, 6, 23), FirstHero: "司马懿", SecondHero: "蔡文姬"},
		{Date: newDate(2023, 7, 21), FirstHero: "曹仁", SecondHero: "夏侯惇"},
		{Date: newDate(2023, 8, 18), FirstHero: "曹丕", SecondHero: "郭嘉"},
		{Date: newDate(2023, 9, 15), FirstHero: "张春华", SecondHero: "张辽"},
		{Date: newDate(2023, 10, 13), FirstHero: "蔡文姬", SecondHero: "夏侯惇"},
		{Date: newDate(2023, 11, 10), FirstHero: "曹丕", SecondHero: "司马懿"},
	},
	"shu": {
		{Date: newDate(2023, 2, 3), FirstHero: "赵云", SecondHero: "诸葛亮"},
		{Date: newDate(2023, 3, 3), FirstHero: "徐庶", SecondHero: "马岱"},
		{Date: newDate(2023, 3, 31), FirstHero: "刘备", SecondHero: "关羽"},
		{Date: newDate(2023, 4, 28), FirstHero: "庞统", SecondHero: "张飞"},
		{Date: newDate(2023, 5, 26), FirstHero: "诸葛亮", SecondHero: "赵云"},
		{Date: newDate(2023, 6, 23), FirstHero: "马岱", SecondHero: "徐庶"},
		{Date: newDate(2023, 7, 21), FirstHero: "关羽", SecondHero: "庞统"},
		{Date: newDate(2023, 8, 18), FirstHero: "张飞", SecondHero: "赵云"},
		{Date: newDate(2023, 9, 15), FirstHero: "诸葛亮", SecondHero: "马岱"},
		{Date: newDate(2023, 10, 13), FirstHero: "关羽", SecondHero: "徐庶"},
		{Date: newDate(2023, 11, 10), FirstHero: "刘备", SecondHero: "庞统"},
	},
	"wu": {
		{Date: newDate(2023, 2, 3), FirstHero: "孙权", SecondHero: "大乔"},
		{Date: newDate(2023, 3, 3), FirstHero: "孙尚香", SecondHero: "小乔"},
		{Date: newDate(2023, 3, 31), FirstHero: "孙坚", SecondHero: "周瑜"},
		{Date: newDate(2023, 4, 28), FirstHero: "孙策", SecondHero: "陆逊"},
		{Date: newDate(2023, 5, 26), FirstHero: "黄盖", SecondHero: "大乔"},
		{Date: newDate(2023, 6, 23), FirstHero: "孙坚", SecondHero: "孙尚香"},
		{Date: newDate(2023, 7, 21), FirstHero: "小乔", SecondHero: "黄盖"},
		{Date: newDate(2023, 8, 18), FirstHero: "孙策", SecondHero: "周瑜"},
		{Date: newDate(2023, 9, 15), FirstHero: "孙权", SecondHero: "陆逊"},
		{Date: newDate(2023, 10, 13), FirstHero: "孙尚香", SecondHero: "大乔"},
		{Date: newDate(2023, 11, 10), FirstHero: "孙坚", SecondHero: "小乔"},
	},
	"qun": {
		{Date: newDate(2023, 2, 3), FirstHero: "吕布", SecondHero: "左慈"},
		{Date: newDate(2023, 3, 3), FirstHero: "貂蝉", SecondHero: "李儒"},
		{Date: newDate(2023, 3, 31), FirstHero: "董卓", SecondHero: "于吉"},
		{Date: newDate(2023, 4, 28), FirstHero: "左慈", SecondHero: "祝融"},
		{Date: newDate(2023, 5, 26), FirstHero: "吕布", SecondHero: "貂蝉"},
		{Date: newDate(2023, 6, 23), FirstHero: "于吉", SecondHero: "董卓"},
		{Date: newDate(2023, 7, 21), FirstHero: "袁绍", SecondHero: "李儒"},
		{Date: newDate(2023, 8, 18), FirstHero: "左慈", SecondHero: "祝融"},
		{Date: newDate(2023, 9, 15), FirstHero: "SP华佗", SecondHero: "董卓"},
		{Date: newDate(2023, 10, 13), FirstHero: "貂蝉", SecondHero: "于吉"},
		{Date: newDate(2023, 11, 10), FirstHero: "吕布", SecondHero: "sp华佗"},
	},
}
