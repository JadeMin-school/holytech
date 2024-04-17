package foodMenu

import (
	json "encoding/json"
	os "os"
)




// 캐시 여부 확인
func isCached() bool {
	_, err := os.Stat("./cache.json")
	return err == nil
}

// 캐시에 저장
func saveCache(weekMenu []*Menu) {
	file, err := os.Create("./cache.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(&weekMenu)
}

// 캐시에서 불러오기
func loadCache() []*Menu {
	weekMenu := []*Menu{}

	file, err := os.Open("./cache.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.Decode(&weekMenu)

	return weekMenu
}