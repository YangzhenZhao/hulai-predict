package main

type guaranteedHeros struct {
	Data      string `json:"data"`
	Date      string `json:"date"`
	IsPredict bool   `json:"isPredict"`
	Type      string `json:"type"`
}

type GuaranteedHeros []guaranteedHeros

func (s GuaranteedHeros) Len() int {
	return len(s)
}

func (s GuaranteedHeros) Less(i, j int) bool {
	return s[i].Date > s[j].Date
}

func (s GuaranteedHeros) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
