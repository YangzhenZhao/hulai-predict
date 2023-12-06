package main

import "time"

func validteDate(lastDate time.Time) bool {
	location, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(location)
	if now.Before(lastDate) {
		return false
	}
	nextDate := lastDate.Add(4 * 7 * 24 * time.Hour)
	nextAllowUpdateTime := nextDate.Add((-2*24 + 17) * time.Hour)
	if now.Before(nextAllowUpdateTime) {
		return false
	}
	return now.Before(nextAllowUpdateTime.Add(7 * 24 * time.Hour))
}
