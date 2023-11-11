package main

import (
	"log"
	"sort"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func genGaochouList(country string) []guaranteedHeros {
	knownData := contryDataMap[country]
	predictHeros := genPredictGaochouList(country, knownData)
	return combineGaochouGuaranteed(knownData.GaochouHistory, predictHeros)
}

func genPredictGaochouList(country string, data countryData) []gaochouPredictHeros {
	var predictRes []gaochouPredictHeros
	currentDate := data.GaochouHistory[len(data.GaochouHistory)-1].Date
	var historyHeroList []HeroList
	for _, item := range data.GaochouHistory {
		historyHeroList = append(historyHeroList, HeroList{
			FirstHero:  item.FirstHero,
			SecondHero: item.SecondHero,
		})
	}
	for i := 0; i < 6; i++ {
		currentDate = currentDate.Add(4 * 7 * 24 * time.Hour)
		firstHero, secondHero := predictNextHeros(data.Heros, zhugongMap[country], historyHeroList)
		predictRes = append(predictRes, gaochouPredictHeros{
			FirstHero:  firstHero,
			SecondHero: secondHero,
			Date:       currentDate,
		})
		historyHeroList = append(historyHeroList, HeroList{
			FirstHero:  firstHero,
			SecondHero: secondHero,
		})
	}
	return predictRes

}

func predictNextHeros(countryHeros []string, zhugong string, historyHeroList []HeroList) (string, string) {
	var historyHeros []string
	for i := 0; i < (len(countryHeros)+1)/2-1; i++ {
		historyHeros = append(historyHeros, historyHeroList[len(historyHeroList)-1-i].FirstHero)
		historyHeros = append(historyHeros, historyHeroList[len(historyHeroList)-1-i].SecondHero)
	}
	nocontainHeros := notContainHeros(countryHeros, historyHeros)
	log.Printf("nocontainHeros: %s\n", spew.Sdump(nocontainHeros))
	if len(nocontainHeros) == 2 {
		return nocontainHeros[0], nocontainHeros[1]
	}
	if len(nocontainHeros) == 1 {
		predictFirst, _ := predictNextHerosByTwoRound(len(countryHeros), zhugong, historyHeroList, nocontainHeros[0])
		return predictFirst, nocontainHeros[0]
	}
	return predictNextHerosByTwoRound(len(countryHeros), zhugong, historyHeroList, "")
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
		return s[i].LastPostion < s[j].LastPostion
	}
	return s[i].Count < s[j].Count
}

func (s Metrics) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func predictNextHerosByTwoRound(countryHerosLen int, zhugong string, historyHeroList []HeroList, excludeHero string) (string, string) {
	metricMap := map[string]heroHistoryMetric{}
	for i := 0; i < countryHerosLen-1; i++ {
		heroList := historyHeroList[len(historyHeroList)-countryHerosLen+i+1]
		if heroList.FirstHero != excludeHero && heroList.FirstHero != zhugong {
			count := 1
			if existMetric, ok := metricMap[heroList.FirstHero]; ok {
				count = existMetric.Count + 1
			}
			metricMap[heroList.FirstHero] = heroHistoryMetric{
				Hero:        heroList.FirstHero,
				Count:       count,
				LastPostion: i * 2,
			}
		}
		if heroList.SecondHero != excludeHero && heroList.SecondHero != zhugong {
			count := 1
			if existMetric, ok := metricMap[heroList.SecondHero]; ok {
				count = existMetric.Count + 1
			}
			metricMap[heroList.SecondHero] = heroHistoryMetric{
				Hero:        heroList.SecondHero,
				Count:       count,
				LastPostion: i*2 + 1,
			}
		}
	}
	var metricList []heroHistoryMetric
	for _, value := range metricMap {
		metricList = append(metricList, value)
	}
	sort.Sort(Metrics(metricList))
	// log.Printf("metricList: %s\n", spew.Sdump(metricList))
	return metricList[0].Hero, metricList[1].Hero
}

func combineGaochouGuaranteed(history []gaochouHeros, predict []gaochouPredictHeros) []guaranteedHeros {
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
