package main

import "time"

type gaochouHeros struct {
	FirstHero  string
	SecondHero string
	Date       time.Time
}

type HeroList struct {
	FirstHero  string
	SecondHero string
}

type gaochouPredictHeros struct {
	FirstHero  string
	SecondHero string
	Date       time.Time
}

type angeHeros struct {
	FirstHero  string
	SecondHero string
	Date       time.Time
}

type angePredictHeros struct {
	FirstHero  string
	SecondHero string
	Date       time.Time
}

type countryData struct {
	Zhugong             string
	LastAngeZhugongDate time.Time
	Heros               []string
	GaochouHistory      []gaochouHeros
	AngeHistory         []angeHeros
}

func newDate(year int, month time.Month, day int) time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Date(year, month, day, 0, 0, 0, 0, location)
}

var zhugongMap = map[string]string{
	"wei": "曹操",
	"shu": "刘备",
	"wu":  "孙权",
	"qun": "袁绍",
}

var contryDataMap = map[string]countryData{
	"wei": {
		LastAngeZhugongDate: newDate(2023, 5, 26),
		Heros:               []string{"曹丕", "张春华", "蔡文姬", "曹仁", "司马懿", "夏侯惇", "张辽", "郭嘉"},
		GaochouHistory: []gaochouHeros{
			{Date: newDate(2023, 2, 17), FirstHero: "郭嘉", SecondHero: "夏侯惇"},
			{Date: newDate(2023, 3, 18), FirstHero: "蔡文姬", SecondHero: "张春华"},
			{Date: newDate(2023, 4, 14), FirstHero: "张辽", SecondHero: "司马懿"},
			{Date: newDate(2023, 5, 12), FirstHero: "曹丕", SecondHero: "夏侯惇"},
			{Date: newDate(2023, 6, 9), FirstHero: "曹仁", SecondHero: "郭嘉"},
			{Date: newDate(2023, 7, 7), FirstHero: "曹丕", SecondHero: "蔡文姬"},
			{Date: newDate(2023, 8, 4), FirstHero: "张春华", SecondHero: "张辽"},
			{Date: newDate(2023, 9, 1), FirstHero: "司马懿", SecondHero: "夏侯惇"},
			{Date: newDate(2023, 9, 29), FirstHero: "曹仁", SecondHero: "曹丕"},
			{Date: newDate(2023, 10, 27), FirstHero: "郭嘉", SecondHero: "蔡文姬"},
		},
		AngeHistory: []angeHeros{
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
		},
	},
	"shu": {
		LastAngeZhugongDate: newDate(2023, 3, 31),
		Heros:               []string{"徐庶", "马岱", "庞统", "诸葛亮", "张飞", "赵云", "关羽"},
		GaochouHistory: []gaochouHeros{
			{Date: newDate(2023, 2, 17), FirstHero: "关羽", SecondHero: "庞统"},
			{Date: newDate(2023, 3, 18), FirstHero: "张飞", SecondHero: "赵云"},
			{Date: newDate(2023, 4, 14), FirstHero: "诸葛亮", SecondHero: "徐庶"},
			{Date: newDate(2023, 5, 12), FirstHero: "马岱", SecondHero: "庞统"},
			{Date: newDate(2023, 6, 9), FirstHero: "关羽", SecondHero: "张飞"},
			{Date: newDate(2023, 7, 7), FirstHero: "诸葛亮", SecondHero: "赵云"},
			{Date: newDate(2023, 8, 4), FirstHero: "徐庶", SecondHero: "马岱"},
			{Date: newDate(2023, 9, 1), FirstHero: "庞统", SecondHero: "关羽"},
			{Date: newDate(2023, 9, 29), FirstHero: "赵云", SecondHero: "张飞"},
			{Date: newDate(2023, 10, 27), FirstHero: "诸葛亮", SecondHero: "徐庶"},
		},
		AngeHistory: []angeHeros{
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
		},
	},
	"wu": {
		LastAngeZhugongDate: newDate(2023, 9, 15),
		Heros:               []string{"孙尚香", "陆逊", "大乔", "孙坚", "黄盖", "周瑜", "孙策", "小乔"},
		GaochouHistory: []gaochouHeros{
			{Date: newDate(2023, 2, 17), FirstHero: "周瑜", SecondHero: "黄盖"},
			{Date: newDate(2023, 3, 18), FirstHero: "孙策", SecondHero: "陆逊"},
			{Date: newDate(2023, 4, 14), FirstHero: "孙尚香", SecondHero: "小乔"},
			{Date: newDate(2023, 5, 12), FirstHero: "孙坚", SecondHero: "大乔"},
			{Date: newDate(2023, 6, 9), FirstHero: "周瑜", SecondHero: "黄盖"},
			{Date: newDate(2023, 7, 7), FirstHero: "陆逊", SecondHero: "孙策"},
			{Date: newDate(2023, 8, 4), FirstHero: "孙尚香", SecondHero: "孙坚"},
			{Date: newDate(2023, 9, 1), FirstHero: "大乔", SecondHero: "小乔"},
			{Date: newDate(2023, 9, 29), FirstHero: "周瑜", SecondHero: "黄盖"},
			{Date: newDate(2023, 10, 27), FirstHero: "孙策", SecondHero: "陆逊"},
		},
		AngeHistory: []angeHeros{
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
		},
	},
	"qun": {
		LastAngeZhugongDate: newDate(2023, 7, 21),
		Heros:               []string{"于吉", "SP华佗", "左慈", "董卓", "李儒", "貂蝉", "祝融", "吕布"},
		GaochouHistory: []gaochouHeros{
			{Date: newDate(2023, 2, 17), FirstHero: "董卓", SecondHero: "于吉"},
			{Date: newDate(2023, 3, 18), FirstHero: "吕布", SecondHero: "左慈"},
			{Date: newDate(2023, 4, 14), FirstHero: "貂蝉", SecondHero: "李儒"},
			{Date: newDate(2023, 5, 12), FirstHero: "董卓", SecondHero: "于吉"},
			{Date: newDate(2023, 6, 9), FirstHero: "吕布", SecondHero: "祝融"},
			{Date: newDate(2023, 7, 7), FirstHero: "左慈", SecondHero: "貂蝉"},
			{Date: newDate(2023, 8, 4), FirstHero: "董卓", SecondHero: "吕布"},
			{Date: newDate(2023, 9, 1), FirstHero: "SP华佗", SecondHero: "李儒"},
			{Date: newDate(2023, 9, 29), FirstHero: "于吉", SecondHero: "祝融"},
			{Date: newDate(2023, 10, 27), FirstHero: "貂蝉", SecondHero: "左慈"},
		},
		AngeHistory: []angeHeros{
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
		},
	},
}
