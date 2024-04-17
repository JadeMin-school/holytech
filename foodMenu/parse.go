package foodMenu

import (
	http "net/http"
	regexp "regexp"

	goquery "github.com/PuerkitoBio/goquery"
)

type MenuTable struct {
	// 아침
	Breakfast string `json:"breakfast"`
	// 점심
	Lunch string `json:"lunch"`
	// 저녁
	Dinner string `json:"dinner"`
}
type Menu struct {
	// 학식 날짜
	Date string `json:"date"`
	// 학식 메뉴
	Table *MenuTable `json:"table"`
}



// 이번주 메뉴 구하기 (캐싱)
func getThisWeek() []*Menu {
	var weekMenu []*Menu

	if isCached() {
		weekMenu = loadCache()
	} else {
		weekMenu = parseThisWeek()
		saveCache(weekMenu)
	}

	return weekMenu
}

// 이번주 메뉴 구하기 (파싱만)
func parseThisWeek() []*Menu {
	result := []*Menu{}
	re := regexp.MustCompile(`\s{2,}|\n+|document.write\(getDay2\('([0-9]{4}-[0-9]{2}-[0-9]{2}).*`)

	response, err := http.Get("https://www.kopo.ac.kr/incheon/content.do?menu=6893")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	html, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}

	tr := html.Find(".tbl_table.menu > tbody > tr")
	tr.Each(func(_ int, s *goquery.Selection) {
		td := s.Find("td")

		result = append(result, &Menu{
			Date: re.ReplaceAllString(td.Eq(0).Text(), "$1"),
			Table: &MenuTable{
				Breakfast: re.ReplaceAllString(td.Eq(1).Text(), ""),
				Lunch: re.ReplaceAllString(td.Eq(2).Text(), ""),
				Dinner: re.ReplaceAllString(td.Eq(3).Text(), ""),
			},
		})
	})

	return result
}