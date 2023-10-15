package main

import (
	"fmt"
	xmlparcer "lessons_hw/pkg/xmlParser"
	"lessons_hw/pkg/xmlParser/dates"
	"lessons_hw/pkg/xmlParser/logic"
)

func main() {
	dates := dates.GenerateDates()
	allPeriodReports := xmlparcer.SendAndParse(dates)
	result := logic.GetMaxMinAVGValute(allPeriodReports)

	fmt.Println("Max valute name is: ", result.MaxValuteName, " cource ", result.MaxValuteValue, " date: ", result.MaxDate)
	fmt.Println("Min valute name is: ", result.MinValuteName, " cource ", result.MinValuteValue, " date: ", result.MinDate)
	fmt.Println("Avg ruble cource is ", result.Avg)
}
