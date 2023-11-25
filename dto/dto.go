package dto

import "time"

type GaochouHeros struct {
	FirstHero  string
	SecondHero string
	Date       time.Time
}

type AngeHeros struct {
	FirstHero  string
	SecondHero string
	Date       time.Time
}

type CountryData struct {
	AngeHistory []AngeHeros
}
