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

type AppendNewGaochouHisotryRecordReq struct {
	Date    time.Time
	WeiList []string
	ShuList []string
	WuList  []string
	QunList []string
}

type AppendNewAngeHisotryRecordReq struct {
	Date    time.Time
	WeiList []string
	ShuList []string
	WuList  []string
	QunList []string
}

type UploadGaoChouReq struct {
	Password string
	Content  string
}

type UploadAngeReq struct {
	Password string
	Content  string
}
