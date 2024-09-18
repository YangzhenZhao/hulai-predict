package main

import (
	"strings"
	"time"

	"github.com/YangzhenZhao/hulai-predict/dto"
	"github.com/YangzhenZhao/hulai-predict/storage"
)

func genGaochouList(country string) []guaranteedHeros {
	gaochouHistory := storage.GetGaochouHistory(country)
	predictHeros := genPredictGaochouList(country, gaochouHistory)
	return combineGaochouGuaranteed(gaochouHistory, predictHeros)
}

func genPredictGaochouList(country string, history []dto.GaochouHeros) []gaochouPredictHeros {
	countryHeros := countryHerosMap[country]
	var predictRes []gaochouPredictHeros
	currentDate := history[len(history)-1].Date
	var historyHeroList [][]string
	for _, item := range history {
		historyHeroList = append(historyHeroList, item.Heros)
	}
	for i := 0; i < 6; i++ {
		currentDate = currentDate.Add(4 * 7 * 24 * time.Hour)
		heros := predictNextAngeHeros(
			countryHeros, zhugongMap[country], historyHeroList, nil, BaodiHeroCnt, "高抽",
		)
		predictRes = append(predictRes, gaochouPredictHeros{
			Heros: heros,
			Date:  currentDate,
		})
		historyHeroList = append(historyHeroList, heros)
	}
	return predictRes

}

type heroHistoryMetric struct {
	Hero        string
	Count       int
	LastPostion int
}

type Metrics []heroHistoryMetric

func (s Metrics) Len() int {
	return len(s)
}

func (s Metrics) Less(i, j int) bool {
	if s[i].Count == s[j].Count {
		return s[i].LastPostion > s[j].LastPostion
	}
	return s[i].Count < s[j].Count
}

func (s Metrics) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func combineGaochouGuaranteed(history []dto.GaochouHeros, predict []gaochouPredictHeros) []guaranteedHeros {
	var res []guaranteedHeros
	for i := len(predict) - 1; i >= 0; i-- {
		res = append(res, guaranteedHeros{
			Data:      strings.Join(predict[i].Heros, " "),
			Date:      dateStr(predict[i].Date),
			IsPredict: true,
			Type:      "高抽",
		})
	}
	for i := len(history) - 1; i >= len(history)-6; i-- {
		res = append(res, guaranteedHeros{
			Data:      strings.Join(history[i].Heros, " "),
			Date:      dateStr(history[i].Date),
			IsPredict: false,
			Type:      "高抽",
		})
	}
	return res
}
