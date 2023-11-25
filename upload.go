package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
	"time"

	"github.com/YangzhenZhao/hulai-predict/dto"
	"github.com/YangzhenZhao/hulai-predict/storage"
	"github.com/gin-gonic/gin"
)

var tmpPassword = "cai"

func unmarshalRequest(c *gin.Context, v any) error {
	if c == nil || c.Request == nil {
		return errors.New("invalid request")
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}

func uploadGaochou(c *gin.Context) {
	req := &dto.UploadGaoChouReq{}
	unmarshalRequest(c, req)
	if req.Password != tmpPassword {
		c.JSON(400, "密码错误")
		return
	}
	weiHisoty := storage.GetGaochouHistory("wei")
	lastDate := weiHisoty[len(weiHisoty)-1].Date
	newDate := lastDate.Add(4 * 7 * 24 * time.Hour)
	content := formatContent(req.Content)

	weiList := getCountryHerosByContent(content, "wei")
	shuList := getCountryHerosByContent(content, "shu")
	wuList := getCountryHerosByContent(content, "wu")
	qunList := getCountryHerosByContent(content, "qun")
	if len(weiList) != 2 || len(shuList) != 2 || len(wuList) != 2 || len(qunList) != 2 {
		c.JSON(400, "解析失败!")
		return
	}
	storage.AppendNewGaochouHisotryRecord(dto.AppendNewGaochouHisotryRecordReq{
		Date:    newDate,
		WeiList: weiList,
		ShuList: shuList,
		WuList:  wuList,
		QunList: qunList,
	})
	c.JSON(200, "")
}

func formatContent(content string) string {
	return strings.Join(strings.Split(content, " "), "")
}

func getCountryHerosByContent(content string, country string) []string {
	var res []string
	for _, hero := range countryHerosMap[country] {
		if strings.Contains(content, hero) {
			res = append(res, hero)
		}
	}
	return res
}
