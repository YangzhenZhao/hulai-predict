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

var gaoChouHistoryKey = "gaoChouHistory"
var gaochouHistoryCacheMap sync.Map
var gaochouFilePath = "/home/ubuntu/hulai-predict/storage/csv_files/gaochou_history.csv"

func GetGaochouHistory(country string) []dto.GaochouHeros {
	return getAllGaochouHistory()[country]
}

func getAllGaochouHistory() map[string][]dto.GaochouHeros {
	cache, _ := gaochouHistoryCacheMap.Load(gaoChouHistoryKey)
	return cache.(map[string][]dto.GaochouHeros)
}

func AppendNewGaochouHisotryRecord(req dto.AppendNewGaochouHisotryRecordReq) {
	allHistory := getAllGaochouHistory()
	allHistory["wei"] = append(allHistory["wei"], dto.GaochouHeros{
		Date:  req.Date,
		Heros: req.WeiList,
	})
	allHistory["shu"] = append(allHistory["shu"], dto.GaochouHeros{
		Date:  req.Date,
		Heros: req.ShuList,
	})
	allHistory["wu"] = append(allHistory["wu"], dto.GaochouHeros{
		Date:  req.Date,
		Heros: req.WuList,
	})
	allHistory["qun"] = append(allHistory["qun"], dto.GaochouHeros{
		Date:  req.Date,
		Heros: req.QunList,
	})
	gaochouHistoryCacheMap.Store(gaoChouHistoryKey, allHistory)
	gaochouCSVAppendRecord(req)
}

func gaochouCSVAppendRecord(req dto.AppendNewGaochouHisotryRecordReq) {
	csvFile, _ := os.OpenFile(gaochouFilePath, os.O_APPEND|os.O_RDWR, os.ModeAppend)
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

func initGaochouHisotyMap() {
	var gaoChouHistoryMap = map[string][]dto.GaochouHeros{
		"wei": {},
		"shu": {},
		"wu":  {},
		"qun": {},
	}
	csvFile, _ := os.OpenFile(gaochouFilePath, os.O_RDWR, os.ModePerm)
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
		gaoChouHistoryMap["wei"] = append(gaoChouHistoryMap["wei"], dto.GaochouHeros{
			Date: date, Heros: weiHeroList,
		})
		gaoChouHistoryMap["shu"] = append(gaoChouHistoryMap["shu"], dto.GaochouHeros{
			Date: date, Heros: shuHeroList,
		})
		gaoChouHistoryMap["wu"] = append(gaoChouHistoryMap["wu"], dto.GaochouHeros{
			Date: date, Heros: wuHeroList,
		})
		gaoChouHistoryMap["qun"] = append(gaoChouHistoryMap["qun"], dto.GaochouHeros{
			Date: date, Heros: qunHeroList,
		})
	}
	gaochouHistoryCacheMap.Store(gaoChouHistoryKey, gaoChouHistoryMap)
}
