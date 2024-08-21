package main

import (
	"time"
)

type HeroList struct {
	FirstHero  string
	SecondHero string
}

type gaochouPredictHeros struct {
	FirstHero  string
	SecondHero string
	Date       time.Time
}

type angePredictHeros struct {
	FirstHero  string
	SecondHero string
	Date       time.Time
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
	"shu": newDate(2024, 6, 21),
	"wu":  newDate(2024, 4, 26),
	"qun": newDate(2024, 3, 1),
}

var countryHerosMap = map[string][]string{
	"wei": {"甄姬", "曹丕", "张春华", "蔡文姬", "曹仁", "司马懿", "夏侯惇", "张辽", "郭嘉"},
	"shu": {"姜维", "徐庶", "马岱", "庞统", "诸葛亮", "张飞", "赵云", "关羽"},
	"wu":  {"步练师", "孙尚香", "陆逊", "大乔", "孙坚", "黄盖", "周瑜", "孙策", "小乔"},
	"qun": {"SP吕布", "于吉", "SP华佗", "左慈", "董卓", "李儒", "貂蝉", "祝融", "吕布"},
}

var countryHerosMapWithZhugong = map[string][]string{
	"wei": {"曹操", "甄姬", "曹丕", "张春华", "蔡文姬", "曹仁", "司马懿", "夏侯惇", "张辽", "郭嘉"},
	"shu": {"刘备", "姜维", "徐庶", "马岱", "庞统", "诸葛亮", "张飞", "赵云", "关羽"},
	"wu":  {"步练师", "孙权", "孙尚香", "陆逊", "大乔", "孙坚", "黄盖", "周瑜", "孙策", "小乔"},
	"qun": {"SP吕布", "袁绍", "于吉", "SP华佗", "左慈", "董卓", "李儒", "貂蝉", "祝融", "吕布"},
}
