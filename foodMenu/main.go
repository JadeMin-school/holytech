package foodMenu

import (
	http "net/http"
	regexp "regexp"
	time "time"

	goquery "github.com/PuerkitoBio/goquery"
)

type MenuTable struct {
	// 아침
	Breakfast string
	// 점심
	Lunch string
	// 저녁
	Dinner string
}
type Menu struct {
	// 학식 날짜
	Date string
	// 학식 메뉴
	Table *MenuTable
}



// 이번주 메뉴 구하기
func GetThisWeek() []*Menu {
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

// 오늘 메뉴 구하기
func GetToday() *Menu {
	currentDate := time.Now().Format("2006-01-02")

	weekMenu := GetThisWeek()
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