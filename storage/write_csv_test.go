package storage

// var writeGaoChouHistoryMap = map[string][]dto.GaochouHeros{
// 	"wei": {
// 		{Date: newDate(2023, 2, 17), FirstHero: "郭嘉", SecondHero: "夏侯惇"},
// 		{Date: newDate(2023, 3, 18), FirstHero: "蔡文姬", SecondHero: "张春华"},
// 		{Date: newDate(2023, 4, 14), FirstHero: "张辽", SecondHero: "司马懿"},
// 		{Date: newDate(2023, 5, 12), FirstHero: "曹丕", SecondHero: "夏侯惇"},
// 		{Date: newDate(2023, 6, 9), FirstHero: "曹仁", SecondHero: "郭嘉"},
// 		{Date: newDate(2023, 7, 7), FirstHero: "曹丕", SecondHero: "蔡文姬"},
// 		{Date: newDate(2023, 8, 4), FirstHero: "张春华", SecondHero: "张辽"},
// 		{Date: newDate(2023, 9, 1), FirstHero: "司马懿", SecondHero: "夏侯惇"},
// 		{Date: newDate(2023, 9, 29), FirstHero: "曹仁", SecondHero: "曹丕"},
// 		{Date: newDate(2023, 10, 27), FirstHero: "郭嘉", SecondHero: "蔡文姬"},
// 	},
// 	"shu": {
// 		{Date: newDate(2023, 2, 17), FirstHero: "关羽", SecondHero: "庞统"},
// 		{Date: newDate(2023, 3, 18), FirstHero: "张飞", SecondHero: "赵云"},
// 		{Date: newDate(2023, 4, 14), FirstHero: "诸葛亮", SecondHero: "徐庶"},
// 		{Date: newDate(2023, 5, 12), FirstHero: "马岱", SecondHero: "庞统"},
// 		{Date: newDate(2023, 6, 9), FirstHero: "关羽", SecondHero: "张飞"},
// 		{Date: newDate(2023, 7, 7), FirstHero: "诸葛亮", SecondHero: "赵云"},
// 		{Date: newDate(2023, 8, 4), FirstHero: "徐庶", SecondHero: "马岱"},
// 		{Date: newDate(2023, 9, 1), FirstHero: "庞统", SecondHero: "关羽"},
// 		{Date: newDate(2023, 9, 29), FirstHero: "赵云", SecondHero: "张飞"},
// 		{Date: newDate(2023, 10, 27), FirstHero: "诸葛亮", SecondHero: "徐庶"},
// 	},
// 	"wu": {
// 		{Date: newDate(2023, 2, 17), FirstHero: "周瑜", SecondHero: "黄盖"},
// 		{Date: newDate(2023, 3, 18), FirstHero: "孙策", SecondHero: "陆逊"},
// 		{Date: newDate(2023, 4, 14), FirstHero: "孙尚香", SecondHero: "小乔"},
// 		{Date: newDate(2023, 5, 12), FirstHero: "孙坚", SecondHero: "大乔"},
// 		{Date: newDate(2023, 6, 9), FirstHero: "周瑜", SecondHero: "黄盖"},
// 		{Date: newDate(2023, 7, 7), FirstHero: "陆逊", SecondHero: "孙策"},
// 		{Date: newDate(2023, 8, 4), FirstHero: "孙尚香", SecondHero: "孙坚"},
// 		{Date: newDate(2023, 9, 1), FirstHero: "大乔", SecondHero: "小乔"},
// 		{Date: newDate(2023, 9, 29), FirstHero: "周瑜", SecondHero: "黄盖"},
// 		{Date: newDate(2023, 10, 27), FirstHero: "孙策", SecondHero: "陆逊"},
// 	},
// 	"qun": {
// 		{Date: newDate(2023, 2, 17), FirstHero: "董卓", SecondHero: "于吉"},
// 		{Date: newDate(2023, 3, 18), FirstHero: "吕布", SecondHero: "左慈"},
// 		{Date: newDate(2023, 4, 14), FirstHero: "貂蝉", SecondHero: "李儒"},
// 		{Date: newDate(2023, 5, 12), FirstHero: "董卓", SecondHero: "于吉"},
// 		{Date: newDate(2023, 6, 9), FirstHero: "吕布", SecondHero: "祝融"},
// 		{Date: newDate(2023, 7, 7), FirstHero: "左慈", SecondHero: "貂蝉"},
// 		{Date: newDate(2023, 8, 4), FirstHero: "董卓", SecondHero: "吕布"},
// 		{Date: newDate(2023, 9, 1), FirstHero: "SP华佗", SecondHero: "李儒"},
// 		{Date: newDate(2023, 9, 29), FirstHero: "于吉", SecondHero: "祝融"},
// 		{Date: newDate(2023, 10, 27), FirstHero: "貂蝉", SecondHero: "左慈"},
// 	},
// }

// var writeAngeHistoryMap = map[string][]dto.AngeHeros{
// 	"wei": []dto.AngeHeros{
// 		{Date: newDate(2023, 2, 3), FirstHero: "张辽", SecondHero: "蔡文姬"},
// 		{Date: newDate(2023, 3, 3), FirstHero: "司马懿", SecondHero: "曹仁"},
// 		{Date: newDate(2023, 3, 31), FirstHero: "夏侯惇", SecondHero: "郭嘉"},
// 		{Date: newDate(2023, 4, 28), FirstHero: "曹丕", SecondHero: "张春华"},
// 		{Date: newDate(2023, 5, 26), FirstHero: "曹操", SecondHero: "张辽"},
// 		{Date: newDate(2023, 6, 23), FirstHero: "司马懿", SecondHero: "蔡文姬"},
// 		{Date: newDate(2023, 7, 21), FirstHero: "曹仁", SecondHero: "夏侯惇"},
// 		{Date: newDate(2023, 8, 18), FirstHero: "曹丕", SecondHero: "郭嘉"},
// 		{Date: newDate(2023, 9, 15), FirstHero: "张春华", SecondHero: "张辽"},
// 		{Date: newDate(2023, 10, 13), FirstHero: "蔡文姬", SecondHero: "夏侯惇"},
// 		{Date: newDate(2023, 11, 10), FirstHero: "曹丕", SecondHero: "司马懿"},
// 	},
// 	"shu": {
// 		{Date: newDate(2023, 2, 3), FirstHero: "赵云", SecondHero: "诸葛亮"},
// 		{Date: newDate(2023, 3, 3), FirstHero: "徐庶", SecondHero: "马岱"},
// 		{Date: newDate(2023, 3, 31), FirstHero: "刘备", SecondHero: "关羽"},
// 		{Date: newDate(2023, 4, 28), FirstHero: "庞统", SecondHero: "张飞"},
// 		{Date: newDate(2023, 5, 26), FirstHero: "诸葛亮", SecondHero: "赵云"},
// 		{Date: newDate(2023, 6, 23), FirstHero: "马岱", SecondHero: "徐庶"},
// 		{Date: newDate(2023, 7, 21), FirstHero: "关羽", SecondHero: "庞统"},
// 		{Date: newDate(2023, 8, 18), FirstHero: "张飞", SecondHero: "赵云"},
// 		{Date: newDate(2023, 9, 15), FirstHero: "诸葛亮", SecondHero: "马岱"},
// 		{Date: newDate(2023, 10, 13), FirstHero: "关羽", SecondHero: "徐庶"},
// 		{Date: newDate(2023, 11, 10), FirstHero: "刘备", SecondHero: "庞统"},
// 	},
// 	"wu": {
// 		{Date: newDate(2023, 2, 3), FirstHero: "孙权", SecondHero: "大乔"},
// 		{Date: newDate(2023, 3, 3), FirstHero: "孙尚香", SecondHero: "小乔"},
// 		{Date: newDate(2023, 3, 31), FirstHero: "孙坚", SecondHero: "周瑜"},
// 		{Date: newDate(2023, 4, 28), FirstHero: "孙策", SecondHero: "陆逊"},
// 		{Date: newDate(2023, 5, 26), FirstHero: "黄盖", SecondHero: "大乔"},
// 		{Date: newDate(2023, 6, 23), FirstHero: "孙坚", SecondHero: "孙尚香"},
// 		{Date: newDate(2023, 7, 21), FirstHero: "小乔", SecondHero: "黄盖"},
// 		{Date: newDate(2023, 8, 18), FirstHero: "孙策", SecondHero: "周瑜"},
// 		{Date: newDate(2023, 9, 15), FirstHero: "孙权", SecondHero: "陆逊"},
// 		{Date: newDate(2023, 10, 13), FirstHero: "孙尚香", SecondHero: "大乔"},
// 		{Date: newDate(2023, 11, 10), FirstHero: "孙坚", SecondHero: "小乔"},
// 	},
// 	"qun": {
// 		{Date: newDate(2023, 2, 3), FirstHero: "吕布", SecondHero: "左慈"},
// 		{Date: newDate(2023, 3, 3), FirstHero: "貂蝉", SecondHero: "李儒"},
// 		{Date: newDate(2023, 3, 31), FirstHero: "董卓", SecondHero: "于吉"},
// 		{Date: newDate(2023, 4, 28), FirstHero: "左慈", SecondHero: "祝融"},
// 		{Date: newDate(2023, 5, 26), FirstHero: "吕布", SecondHero: "貂蝉"},
// 		{Date: newDate(2023, 6, 23), FirstHero: "于吉", SecondHero: "董卓"},
// 		{Date: newDate(2023, 7, 21), FirstHero: "袁绍", SecondHero: "李儒"},
// 		{Date: newDate(2023, 8, 18), FirstHero: "左慈", SecondHero: "祝融"},
// 		{Date: newDate(2023, 9, 15), FirstHero: "SP华佗", SecondHero: "董卓"},
// 		{Date: newDate(2023, 10, 13), FirstHero: "貂蝉", SecondHero: "于吉"},
// 		{Date: newDate(2023, 11, 10), FirstHero: "吕布", SecondHero: "sp华佗"},
// 	},
// }

// func Test_writeGaochouCSV(t *testing.T) {
// 	filePath := "/home/ubuntu/hulai-predict/storage/csv_files/gaochou_history.csv"
// 	csvFile, _ := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
// 	w := csv.NewWriter(csvFile)
// 	w.Write([]string{"year", "month", "day", "wei", "shu", "wu", "qun"})
// 	for idx, item := range writeGaoChouHistoryMap["wei"] {
// 		year := item.Date.Year()
// 		month := item.Date.Month()
// 		day := item.Date.Day()
// 		shu := []string{writeGaoChouHistoryMap["shu"][idx].FirstHero, writeGaoChouHistoryMap["shu"][idx].SecondHero}
// 		wu := []string{writeGaoChouHistoryMap["wu"][idx].FirstHero, writeGaoChouHistoryMap["wu"][idx].SecondHero}
// 		qun := []string{writeGaoChouHistoryMap["qun"][idx].FirstHero, writeGaoChouHistoryMap["qun"][idx].SecondHero}
// 		record := []string{
// 			strconv.Itoa(year),
// 			strconv.Itoa(int(month)),
// 			strconv.Itoa(int(day)),
// 			strings.Join([]string{item.FirstHero, item.SecondHero}, " "),
// 			strings.Join(shu, " "),
// 			strings.Join(wu, " "),
// 			strings.Join(qun, " "),
// 		}
// 		w.Write(record)
// 	}
// 	w.Flush()
// 	if err := w.Error(); err != nil {
// 		fmt.Printf("err:%v", err)
// 	}
// }

// func Test_writeAngeCSV(t *testing.T) {
// 	filePath := "/home/ubuntu/hulai-predict/storage/csv_files/ange_history.csv"
// 	csvFile, _ := os.OpenFile(filePath, os.O_RDWR, os.ModePerm)
// 	w := csv.NewWriter(csvFile)
// 	w.Write([]string{"year", "month", "day", "wei", "shu", "wu", "qun"})
// 	for idx, item := range writeAngeHistoryMap["wei"] {
// 		year := item.Date.Year()
// 		month := item.Date.Month()
// 		day := item.Date.Day()
// 		shu := []string{writeAngeHistoryMap["shu"][idx].FirstHero, writeAngeHistoryMap["shu"][idx].SecondHero}
// 		wu := []string{writeAngeHistoryMap["wu"][idx].FirstHero, writeAngeHistoryMap["wu"][idx].SecondHero}
// 		qun := []string{writeAngeHistoryMap["qun"][idx].FirstHero, writeAngeHistoryMap["qun"][idx].SecondHero}
// 		record := []string{
// 			strconv.Itoa(year),
// 			strconv.Itoa(int(month)),
// 			strconv.Itoa(int(day)),
// 			strings.Join([]string{item.FirstHero, item.SecondHero}, " "),
// 			strings.Join(shu, " "),
// 			strings.Join(wu, " "),
// 			strings.Join(qun, " "),
// 		}
// 		w.Write(record)
// 	}
// 	w.Flush()
// 	if err := w.Error(); err != nil {
// 		fmt.Printf("err:%v", err)
// 	}
// }
