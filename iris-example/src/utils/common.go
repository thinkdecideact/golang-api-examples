package utils

import (
	"encoding/json"
	"math"
	"time"
)

func CalPageCount(rowCount int64, rowCountPerPage int64) int64 {
	if rowCount%rowCountPerPage == 0 {
		return rowCount / rowCountPerPage
	}
	return int64(math.Ceil(float64(rowCount) / float64(rowCountPerPage)))
}

func ConvertStructToMap(structData interface{}) map[string]interface{} {
	var mapData map[string]interface{}
	temp, _ := json.Marshal(structData)
	json.Unmarshal(temp, &mapData)
	return mapData
}

func GetFormattedTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func GetFormattedNowTime() string {
	var ChinaZone = time.FixedZone("CST", 8*3600) // GMT+8
	return time.Now().In(ChinaZone).Format("2006-01-02 15:04:05")
}
