package storage

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/YangzhenZhao/hulai-predict/dto"
)

func GetGaochouHistory(country string) []dto.GaochouHeros {
	return gaoChouHistoryMap[country]
}

var gaochouFilePath = "/home/ubuntu/hulai-predict/storage/csv_files/gaochou_history.csv"

var gaoChouHistoryMap = map[string][]dto.GaochouHeros{
	"wei": {},
	"shu": {},
	"wu":  {},
	"qun": {},
}

func initGaochouHisotyMap() {
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
			Date: date, FirstHero: weiHeroList[0], SecondHero: weiHeroList[1],
		})
		gaoChouHistoryMap["shu"] = append(gaoChouHistoryMap["shu"], dto.GaochouHeros{
			Date: date, FirstHero: shuHeroList[0], SecondHero: shuHeroList[1],
		})
		gaoChouHistoryMap["wu"] = append(gaoChouHistoryMap["wu"], dto.GaochouHeros{
			Date: date, FirstHero: wuHeroList[0], SecondHero: wuHeroList[1],
		})
		gaoChouHistoryMap["qun"] = append(gaoChouHistoryMap["qun"], dto.GaochouHeros{
			Date: date, FirstHero: qunHeroList[0], SecondHero: qunHeroList[1],
		})
	}
}
