package foodMenu

import (
	fmt "fmt"
	time "time"
)

var (
	weekMenu []*Menu
	todayMenu *Menu
)

func init() {
	weekMenu = getThisWeek()
	todayMenu = filterToday(weekMenu)
}



// 캐시 초기화
func hitCache() {
	currentDate := time.Now().Format("2006-01-02")

	if todayMenu.Date != currentDate {
		fmt.Print("메모리에 캐시된 오늘의 메뉴가 현재 날짜와 불일치합니다.\n오늘의 메뉴를 다시 로딩합니다.\n\n")
		todayMenu = filterToday(weekMenu)
	}
	if todayMenu.Table == nil {
		fmt.Print("메모리에 캐시된 이번주 메뉴에서 현재 날짜에 해당하는 메뉴를 찾을 수 없습니다.\n이번주 메뉴를 다시 로딩합니다.\n\n")
		weekMenu = getThisWeek()
		todayMenu = filterToday(weekMenu)
	}
}

// 한 주 메뉴 리스트에서 오늘 메뉴만 필터링
func filterToday(weekMenu []*Menu) *Menu {
	currentDate := time.Now().Format("2006-01-02")

	for _, menu := range weekMenu {
		if menu.Date == currentDate {
			return menu
		}
	}

	return &Menu{
		Date: currentDate,
		Table: nil,
	}
}

// 이번주 메뉴 구하기
func GetThisWeek() []*Menu {
	hitCache()
	
	return weekMenu
}

// 오늘 메뉴 구하기
func GetToday() *Menu {
	hitCache()

	return todayMenu
}