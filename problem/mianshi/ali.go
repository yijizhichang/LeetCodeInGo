package mianshi

import (
	"encoding/json"
	"strings"
)

type AreaResource struct {
	Area     string
	Splitter string
	count    int64
}
// dfs
func getFormattedJSONByResource(areas []AreaResource) string{
	area2count := make(map[string]interface{})
	for _, area := range areas {
		tmps := strings.Split(area.Area, area.Splitter)
		getRetM(tmps, len(tmps), len(tmps), area.count, area2count)
	}
	res, err := json.Marshal(area2count)
	if err != nil {
		return ""
	}
	return string(res)
}

func getRetM(tmps []string, last, length int, count int64, now interface{}) {
	key := tmps[length - last]
	if newNow, ok := now.(map[string]interface{}); ok {
		if last == 1 {
			if value, ok := newNow[key]; ok {
				if value64, ok := value.(int64); ok {
					newNow[key] = value64 + count
					return
				}
			} else {
				newNow[key] = count
			}
		} else {
			if _, ok := newNow[key]; ok {
				// 如果存在当前key的map
				getRetM(tmps, last-1, length, count, newNow[key])
			} else {
				new := make(map[string]interface{})
				getRetM(tmps, last-1, length, count, new)
				newNow[key] = new
			}
		}
		now = newNow
	}

}

// bfs
type ProvinceService interface {
	GetFormattedJSONByResource(areas []AreaResource) string
}

type ProvinceServiceIns struct {
}

func (ps ProvinceServiceIns) GetFormattedJSONByResource(areas []AreaResource) string {
	tmpMap:=make(map[string]interface{})
	for _,ar:=range areas{
		var a map[string]interface{}
		a = tmpMap
		areaArr := strings.Split(ar.Area,ar.Splitter)
		for i,area:=range areaArr{
			if v,ok:=a[area];ok{
				a = v.(map[string]interface{})
			}else{
				if i==len(areaArr)-1{
					a[area] = ar.count
				}else{
					a[area] = make(map[string]interface{})
					a=a[area].(map[string]interface{})
				}
			}
		}
	}
	res,_:=json.Marshal(tmpMap)
	return string(res)
}

