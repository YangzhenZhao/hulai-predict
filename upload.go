package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"sort"
	"strings"
	"time"

	"github.com/YangzhenZhao/hulai-predict/dto"
	"github.com/YangzhenZhao/hulai-predict/storage"
	"github.com/gin-gonic/gin"
)

var tmpPassword = "cai"

const BaodiHeroCnt = 3

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
	// if !validteDate(lastDate) {
	// 	c.JSON(400, "当前不可更新高抽")
	// 	return
	// }
	newDate := lastDate.Add(4 * 7 * 24 * time.Hour)
	content := formatContent(req.Content)

	weiList := getCountryHerosByContent(content, "wei")
	shuList := getCountryHerosByContent(content, "shu")
	wuList := getCountryHerosByContent(content, "wu")
	qunList := getCountryHerosByContent(content, "qun")
	if len(weiList) != BaodiHeroCnt || len(shuList) != BaodiHeroCnt || len(wuList) != BaodiHeroCnt || len(qunList) != BaodiHeroCnt {
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

func uploadAnge(c *gin.Context) {
	req := &dto.UploadAngeReq{}
	unmarshalRequest(c, req)
	if req.Password != tmpPassword {
		c.JSON(400, "密码错误")
		return
	}
	weiHisoty := storage.GetAngeHistory("wei")
	lastDate := weiHisoty[len(weiHisoty)-1].Date
	// if !validteDate(lastDate) {
	// 	c.JSON(400, "当前不可更新暗格")
	// 	return
	// }
	newDate := lastDate.Add(4 * 7 * 24 * time.Hour)
	content := formatContent(req.Content)

	weiList := getCountryHerosByContent(content, "wei")
	shuList := getCountryHerosByContent(content, "shu")
	wuList := getCountryHerosByContent(content, "wu")
	qunList := getCountryHerosByContent(content, "qun")
	if len(weiList) != BaodiHeroCnt || len(shuList) != BaodiHeroCnt || len(wuList) != BaodiHeroCnt || len(qunList) != BaodiHeroCnt {
		c.JSON(400, "解析失败!")
		return
	}
	storage.AppendNewAngeHisotryRecord(dto.AppendNewAngeHisotryRecordReq{
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
	var idxList []int
	for _, hero := range countryHerosMapWithZhugong[country] {
		if index := strings.Index(content, hero); index != -1 {
			res = append(res, hero)
			idxList = append(idxList, index)
			if len(res) == BaodiHeroCnt {
				break
			}
		}
	}
	if len(res) != BaodiHeroCnt {
		return nil
	}

	sort.Slice(res, func(i, j int) bool {
		return idxList[i] < idxList[j]
	})
	return res
}
