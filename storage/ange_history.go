package storage

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/YangzhenZhao/hulai-predict/dto"
)

func newDate(year int, month time.Month, day int) time.Time {
	location, _ := time.LoadLocation("Asia/Shanghai")
	return time.Date(year, month, day, 0, 0, 0, 0, location)
}

var angeHistoryKey = "angeHistory"
var angeHistoryCacheMap sync.Map
var angeFilePath = "/home/ubuntu/hulai-predict/storage/csv_files/ange_history.csv"

func GetAngeHistory(country string) []dto.AngeHeros {
	return getAllAngeHistory()[country]
}

func getAllAngeHistory() map[string][]dto.AngeHeros {
	cache, _ := angeHistoryCacheMap.Load(angeHistoryKey)
	return cache.(map[string][]dto.AngeHeros)
}

func AppendNewAngeHisotryRecord(req dto.AppendNewAngeHisotryRecordReq) {
	allHistory := getAllAngeHistory()
	allHistory["wei"] = append(allHistory["wei"], dto.AngeHeros{
		Date:  req.Date,
		Heros: req.WeiList,
	})
	allHistory["shu"] = append(allHistory["shu"], dto.AngeHeros{
		Date:  req.Date,
		Heros: req.ShuList,
	})
	allHistory["wu"] = append(allHistory["wu"], dto.AngeHeros{
		Date:  req.Date,
		Heros: req.WuList,
	})
	allHistory["qun"] = append(allHistory["qun"], dto.AngeHeros{
		Date:  req.Date,
		Heros: req.QunList,
	})
	angeHistoryCacheMap.Store(angeHistoryKey, allHistory)
	angeCSVAppendRecord(req)
}

func angeCSVAppendRecord(req dto.AppendNewAngeHisotryRecordReq) {
	csvFile, _ := os.OpenFile(angeFilePath, os.O_APPEND|os.O_RDWR, os.ModeAppend)
	w := csv.NewWriter(csvFile)
	record := []string{
		strconv.Itoa(req.Date.Year()),
		strconv.Itoa(int(req.Date.Month())),
		strconv.Itoa(req.Date.Day()),
		strings.Join(req.WeiList, " "),
		strings.Join(req.ShuList, " "),
		strings.Join(req.WuList, " "),
		strings.Join(req.QunList, " "),
	}
	w.Write(record)
	w.Flush()
}

var angeHistoryMap map[string][]dto.AngeHeros

func initAngeHistoryMap() {
	angeHistoryMap = map[string][]dto.AngeHeros{
		"wei": {},
		"shu": {},
		"wu":  {},
		"qun": {},
	}
	csvFile, _ := os.OpenFile(angeFilePath, os.O_RDWR, os.ModePerm)
	r := csv.NewReader(csvFile)
	records, _ := r.ReadAll()
	for _, record := range records[1:] {
		year, _ := strconv.ParseInt(record[0], 10, 32)
		month, _ := strconv.ParseInt(record[1], 10, 32)
		day, _ := strconv.ParseInt(record[2], 10, 32)
		date := newDate(int(year), time.Month(month), int(day))
		weiHeroList := strings.Split(record[3], " ")
		shuHeroList := strings.Split(record[4], " ")
		wuHeroList := strings.Split(record[5], " ")
		qunHeroList := strings.Split(record[6], " ")
		angeHistoryMap["wei"] = append(angeHistoryMap["wei"], dto.AngeHeros{
			Date: date, Heros: weiHeroList,
		})
		angeHistoryMap["shu"] = append(angeHistoryMap["shu"], dto.AngeHeros{
			Date: date, Heros: shuHeroList,
		})
		angeHistoryMap["wu"] = append(angeHistoryMap["wu"], dto.AngeHeros{
			Date: date, Heros: wuHeroList,
		})
		angeHistoryMap["qun"] = append(angeHistoryMap["qun"], dto.AngeHeros{
			Date: date, Heros: qunHeroList,
		})
	}
	angeHistoryCacheMap.Store(angeHistoryKey, angeHistoryMap)
}
