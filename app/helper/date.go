package helper

import (
	"fmt"
	"time"
)

func SubstractDate(startDate string, endDate string) (uint, error) {
	startDateParse, errStartDate := time.Parse("02/01/2006", startDate)
	if errStartDate != nil {
		fmt.Println("err start date parse: " + errStartDate.Error())
		return 0, errStartDate
	}
	endDateParse, errEndDate := time.Parse("02/01/2006", endDate)
	if errEndDate != nil {
		return 0, errEndDate
	}

	duration := endDateParse.Sub(startDateParse)
	return uint(duration.Hours() / 24), nil
}