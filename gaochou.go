package main

import (
	"sort"
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
		historyHeroList = append(historyHeroList, []string{
			item.FirstHero,
			item.SecondHero,
		})
	}
	for i := 0; i < 6; i++ {
		currentDate = currentDate.Add(4 * 7 * 24 * time.Hour)
		firstHero, secondHero := predictNextHeros(countryHeros, zhugongMap[country], historyHeroList)
		predictRes = append(predictRes, gaochouPredictHeros{
			FirstHero:  firstHero,
			SecondHero: secondHero,
			Date:       currentDate,
		})
		historyHeroList = append(historyHeroList, []string{
			firstHero,
			secondHero,
		})
	}
	return predictRes

}

func predictNextHeros(countryHeros []string, zhugong string, historyHeroList [][]string) (string, string) {
	var historyHeros []string
	for i := 0; i < (len(countryHeros)+1)/2-1; i++ {
		historyHeros = append(historyHeros, historyHeroList[len(historyHeroList)-1-i][0])
		historyHeros = append(historyHeros, historyHeroList[len(historyHeroList)-1-i][1])
	}
	nocontainHeros := notContainHeros(countryHeros, historyHeros, nil)
	if len(nocontainHeros) == 2 {
		return nocontainHeros[0], nocontainHeros[1]
	}
	if len(nocontainHeros) == 1 {
		predictFirst, _ := predictNextHerosByTwoRound(countryHeros, len(countryHeros), zhugong, historyHeroList, nocontainHeros[0])
		return predictFirst, nocontainHeros[0]
	}
	return predictNextHerosByTwoRound(countryHeros, len(countryHeros), zhugong, historyHeroList, "")
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

func predictNextHerosByTwoRound(countryHeros []string, countryHerosLen int, zhugong string, historyHeroList [][]string, excludeHero string) (string, string) {
	metricMap := map[string]heroHistoryMetric{}
	for i := 0; i < countryHerosLen-1; i++ {
		heroList := historyHeroList[len(historyHeroList)-countryHerosLen+i+1]
		if heroList[0] != excludeHero && heroList[0] != zhugong {
			count := 1
			if existMetric, ok := metricMap[heroList[0]]; ok {
				count = existMetric.Count + 1
			}
			metricMap[heroList[0]] = heroHistoryMetric{
				Hero:        heroList[0],
				Count:       count,
				LastPostion: i * 2,
			}
		}
		if heroList[1] != excludeHero && heroList[1] != zhugong {
			count := 1
			if existMetric, ok := metricMap[heroList[1]]; ok {
				count = existMetric.Count + 1
			}
			metricMap[heroList[1]] = heroHistoryMetric{
				Hero:        heroList[1],
				Count:       count,
				LastPostion: i*2 + 1,
			}
		}
	}
	var metricList []heroHistoryMetric
	var historyHeros []string
	for _, value := range metricMap {
		historyHeros = append(historyHeros, value.Hero)
		metricList = append(metricList, value)
	}
	nocontainHeros := notContainHeros(countryHeros, historyHeros, []string{excludeHero})
	if len(nocontainHeros) >= 2 {
		return nocontainHeros[0], nocontainHeros[1]
	}
	sort.Sort(Metrics(metricList))
	if len(nocontainHeros) == 1 && nocontainHeros[0] != metricList[0].Hero {
		return nocontainHeros[0], metricList[0].Hero
	}
	// log.Printf("metricList: %s\n", spew.Sdump(metricList))
	return metricList[0].Hero, metricList[1].Hero
}

func combineGaochouGuaranteed(history []dto.GaochouHeros, predict []gaochouPredictHeros) []guaranteedHeros {
	var res []guaranteedHeros
	for i := len(predict) - 1; i >= 0; i-- {
		res = append(res, guaranteedHeros{
			Data:      strings.Join([]string{predict[i].FirstHero, predict[i].SecondHero}, " "),
			Date:      dateStr(predict[i].Date),
			IsPredict: true,
			Type:      "高抽",
		})
	}
	for i := len(history) - 1; i >= len(history)-6; i-- {
		res = append(res, guaranteedHeros{
			Data:      strings.Join([]string{history[i].FirstHero, history[i].SecondHero}, " "),
			Date:      dateStr(history[i].Date),
			IsPredict: false,
			Type:      "高抽",
		})
	}
	return res
}
