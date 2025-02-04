package main

import (
	"time"
)

type gaochouPredictHeros struct {
	Heros []string
	Date  time.Time
}

type angePredictHeros struct {
	Heros []string
	Date  time.Time
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

var lastAngeZhugongDateMap = map[string]time.Time{
	"wei": newDate(2024, 8, 16),
	"shu": newDate(2025, 1, 31),
	"wu":  newDate(2024, 12, 6),
	"qun": newDate(2024, 10, 11),
}

var countryHerosMap = map[string][]string{
	"wei": {"甄姬", "曹丕", "张春华", "蔡文姬", "曹仁", "司马懿", "夏侯惇", "张辽", "郭嘉"},
	"shu": {"SP马超", "姜维", "徐庶", "马岱", "庞统", "诸葛亮", "张飞", "赵云", "关羽"},
	"wu":  {"SP甘宁", "步练师", "孙尚香", "陆逊", "大乔", "孙坚", "黄盖", "周瑜", "孙策", "小乔"},
	"qun": {"SP吕布", "于吉", "SP华佗", "左慈", "董卓", "李儒", "貂蝉", "祝融", "吕布"},
}

var countryHerosMapWithZhugong = map[string][]string{
	"wei": {"曹操", "甄姬", "曹丕", "张春华", "蔡文姬", "曹仁", "司马懿", "夏侯惇", "张辽", "郭嘉"},
	"shu": {"刘备", "SP马超", "姜维", "徐庶", "马岱", "庞统", "诸葛亮", "张飞", "赵云", "关羽"},
	"wu":  {"SP甘宁","步练师", "孙权", "孙尚香", "陆逊", "大乔", "孙坚", "黄盖", "周瑜", "孙策", "小乔"},
	"qun": {"SP吕布", "袁绍", "于吉", "SP华佗", "左慈", "董卓", "李儒", "貂蝉", "祝融", "吕布"},
}
