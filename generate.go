package main

import (
	"slices"
	"sort"
)

func generateRes() [][]guaranteedHeros {
	return [][]guaranteedHeros{
		genSingleCountryRes("shu"),
		genSingleCountryRes("wei"),
		genSingleCountryRes("wu"),
		genSingleCountryRes("qun"),
	}
}

func genSingleCountryRes(country string) []guaranteedHeros {
	res := genAngeList(country)
	res = append(res, genGaochouList(country)...)
	sort.Sort(GuaranteedHeros(res))
	return res
}

func notContainHeros(wholeHeros, partHeros []string, excludeHeros []string) []string {
	partHeroMap := map[string]struct{}{}
	for _, hero := range partHeros {
		partHeroMap[hero] = struct{}{}
	}
	var resHeros []string
	for _, hero := range wholeHeros {
		if slices.Contains(excludeHeros, hero) {
			continue
		}
		if _, ok := partHeroMap[hero]; !ok {
			resHeros = append(resHeros, hero)
		}
	}
	return resHeros
}
