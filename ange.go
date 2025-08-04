package main

import (
	"slices"
	"sort"
	"strings"
	"time"

	"github.com/YangzhenZhao/hulai-predict/dto"
	"github.com/YangzhenZhao/hulai-predict/storage"
)

func genAngeList(country string) []guaranteedHeros {
	history := storage.GetAngeHistory(country)
	predictHeros := genPredictAngeList(country, history)
	return combineAngeGuaranteed(history, predictHeros)
}

func getNextAngeZhugongDate(lastHistoryDate time.Time, country string) time.Time {
	newDate := lastAngeZhugongDateMap[country].Add(8 * 4 * 7 * 24 * time.Hour)
	for {
		if newDate.After(lastHistoryDate) {
			break
		}
		newDate = newDate.Add(8 * 4 * 7 * 24 * time.Hour)
	}
	return newDate
}

func genPredictAngeList(country string, history []dto.AngeHeros) []angePredictHeros {
	countryHeros := countryHerosMap[country]
	var predictRes []angePredictHeros
	currentDate := history[len(history)-1].Date
	nextAngeZhugongDate := getNextAngeZhugongDate(currentDate, country)
	var historyHeroList [][]string
	for _, item := range history {
		historyHeroList = append(historyHeroList, item.Heros)
	}
	for i := 0; i < 6; i++ {
		currentDate = currentDate.Add(4 * 7 * 24 * time.Hour)
		var heros []string
		if currentDate.Equal(nextAngeZhugongDate) {
			heros = append(heros, zhugongMap[country])
		}
		heros = append(
			heros,
			predictNextAngeHeros(
				countryHeros, zhugongMap[country], historyHeroList, heros, BaodiHeroCnt-len(heros), "暗格",
			)...)
		predictRes = append(predictRes, angePredictHeros{
			Heros: heros,
			Date:  currentDate,
		})
		historyHeroList = append(historyHeroList, heros)
	}
	return predictRes
}

func predictNextAngeHeros(countryHeros []string, zhugong string, historyHeroList [][]string, excludeHeros []string, leaveCnt int, predictType string) []string {
	oneRoundRes := predictNextHerosByOneRound(countryHeros, zhugong, historyHeroList, excludeHeros)
	if len(oneRoundRes) == leaveCnt {
		return oneRoundRes
	}
	var res []string
	if len(oneRoundRes) < leaveCnt {
		res = append(res, oneRoundRes...)
	}
	return append(res, predictAngeNextHerosByTwoRound(countryHeros, zhugong, historyHeroList, res, leaveCnt-len(res))...)
}

func predictNextHerosByOneRound(countryHeros []string, zhugong string, historyHeroList [][]string, excludeHeros []string) []string {
	var historyHeros []string
	historyHeros = append(historyHeros, excludeHeros...)
	for i := len(historyHeroList) - 1; i >= 0; i-- {
		if len(historyHeros)+len(historyHeroList[i]) >= len(countryHeros) {
			break
		}
		historyHeros = append(historyHeros, historyHeroList[i]...)
	}
	nocontainHeros := notContainHeros(countryHeros, historyHeros, append(excludeHeros, zhugong))
	return nocontainHeros
}

func predictAngeNextHerosByTwoRound(countryHeros []string, zhugong string, historyHeroList [][]string, excludeHeros []string, leaveCnt int) []string {
	metricMap := map[string]heroHistoryMetric{}
	for _, hero := range countryHeros {
		metricMap[hero] = heroHistoryMetric{Hero: hero}
	}
	heroCnt := 0
	for i := len(historyHeroList) - 1; i >= 0; i-- {
		if heroCnt+len(historyHeroList[i]) >= 2*len(countryHeros) {
			break
		}
		heroList := historyHeroList[i]
		for j := 0; j < len(heroList); j++ {
			metricMap[heroList[j]] = heroHistoryMetric{
				Hero:        heroList[j],
				Count:       metricMap[heroList[j]].Count + 1,
				LastPostion: heroCnt + j,
			}
		}
		heroCnt += len(heroList)
	}
	var metricList []heroHistoryMetric
	for _, value := range metricMap {
		if value.Hero == zhugong || slices.Contains(excludeHeros, value.Hero) {
			continue
		}
		metricList = append(metricList, value)
	}
	sort.Sort(Metrics(metricList))
	var res []string
	for i := 0; i < leaveCnt; i++ {
		res = append(res, metricList[i].Hero)
	}
	return res
}

func combineAngeGuaranteed(history []dto.AngeHeros, predict []angePredictHeros) []guaranteedHeros {
	var res []guaranteedHeros
	for i := len(predict) - 1; i >= 0; i-- {
		res = append(res, guaranteedHeros{
			Data:      strings.Join(predict[i].Heros, " "),
			Date:      dateStr(predict[i].Date),
			IsPredict: true,
			Type:      "暗格",
		})
	}
	for i := len(history) - 1; i >= len(history)-6; i-- {
		res = append(res, guaranteedHeros{
			Data:      strings.Join(history[i].Heros, " "),
			Date:      dateStr(history[i].Date),
			IsPredict: false,
			Type:      "暗格",
		})
	}
	return res
}
