package main

import (
	"log"
	"strings"
	"time"

	"github.com/YangzhenZhao/hulai-predict/dto"
	"github.com/YangzhenZhao/hulai-predict/storage"
	"github.com/davecgh/go-spew/spew"
)

func genAngeList(country string) []guaranteedHeros {
	history := storage.GetAngeHistory(country)
	predictHeros := genPredictAngeList(country, history)
	return combineAngeGuaranteed(history, predictHeros)
}

func genPredictAngeList(country string, history []dto.AngeHeros) []angePredictHeros {
	countryHeros := countryHerosMap[country]
	var predictRes []angePredictHeros
	currentDate := history[len(history)-1].Date
	nextAngeZhugongDate := lastAngeZhugongDateMap[country].Add(8 * 4 * 7 * 24 * time.Hour)
	log.Printf("nextAngeZhugongDate: %s", nextAngeZhugongDate.String())
	var historyHeroList []HeroList
	for _, item := range history {
		historyHeroList = append(historyHeroList, HeroList{
			FirstHero:  item.FirstHero,
			SecondHero: item.SecondHero,
		})
	}
	for i := 0; i < 6; i++ {
		currentDate = currentDate.Add(4 * 7 * 24 * time.Hour)
		var firstHero string
		var secondHero string
		if currentDate.Equal(nextAngeZhugongDate) {
			firstHero = zhugongMap[country]
			secondHero = predictNextSingleHero(countryHeros, zhugongMap[country], historyHeroList)
		} else {
			firstHero, secondHero = predictNextHeros(countryHeros, zhugongMap[country], historyHeroList)
		}
		predictRes = append(predictRes, angePredictHeros{
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

func predictNextSingleHero(countryHeros []string, zhugong string, historyHeroList []HeroList) string {
	var historyHeros []string
	for i := 0; i < (len(countryHeros)+1)/2-1; i++ {
		historyHeros = append(historyHeros, historyHeroList[len(historyHeroList)-1-i].FirstHero)
		historyHeros = append(historyHeros, historyHeroList[len(historyHeroList)-1-i].SecondHero)
	}
	nocontainHeros := notContainHeros(countryHeros, historyHeros)
	log.Printf("nocontainHeros: %s\n", spew.Sdump(nocontainHeros))
	if len(nocontainHeros) == 1 {
		return nocontainHeros[0]
	}
	predictFirst, _ := predictNextHerosByTwoRound(len(countryHeros), zhugong, historyHeroList, "")
	return predictFirst
}

func combineAngeGuaranteed(history []dto.AngeHeros, predict []angePredictHeros) []guaranteedHeros {
	var res []guaranteedHeros
	for i := len(predict) - 1; i >= 0; i-- {
		res = append(res, guaranteedHeros{
			Data:      strings.Join([]string{predict[i].FirstHero, predict[i].SecondHero}, " "),
			Date:      dateStr(predict[i].Date),
			IsPredict: true,
			Type:      "暗格",
		})
	}
	for i := len(history) - 1; i >= len(history)-6; i-- {
		res = append(res, guaranteedHeros{
			Data:      strings.Join([]string{history[i].FirstHero, history[i].SecondHero}, " "),
			Date:      dateStr(history[i].Date),
			IsPredict: false,
			Type:      "暗格",
		})
	}
	return res
}
