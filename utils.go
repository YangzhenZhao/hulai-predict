package main

import (
	"fmt"
	"time"
)

func dateStr(date time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d", date.Year(), date.Month(), date.Day())
}
