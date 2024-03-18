package main

import (
	fmt "fmt"
)
import (
	foodMenu "holytech/foodMenu"
)



func main() {
	menu := foodMenu.GetToday()
	
	if menu.Table == nil {
		fmt.Printf(
			"날짜 : %s\n급식이 없습니다.\n\n",
			menu.Date,
		)
	} else {
		fmt.Printf(
			"날짜 : %s\n아침 : %s\n점심 : %s\n저녁 : %s\n\n",
			menu.Date, menu.Table.Breakfast, menu.Table.Lunch, menu.Table.Dinner,
		)
	}
}