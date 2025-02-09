package main

import "fmt"

const (
	Monday int = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func IsLeapYear(year int) bool {
	return ((year%4 == 0) && (year%100 != 0)) || (year%400 == 0)
}

func DaysInMonth(month, year int) int {
	days := 0

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		days = 31
	case 4, 6, 9, 11:
		days = 30
	case 2:
		if IsLeapYear(year) {
			days = 29
		} else {
			days = 28
		}
	}

	return days
}

func calculateDayOfWeek(refYear, refMonth, refDay, year, month, date int) int {

	numDays := 0

	if refYear > year {
		fmt.Println("TODO: Day in past than reference!!!")
		return 0
	}

	//years from reference
	for y := refYear; y < year; y++ {
		if IsLeapYear(y) {
			numDays += 366
		} else {
			numDays += 365
		}
	}

	//months from reference
	for m := refMonth; m < month; m++ {
		numDays += DaysInMonth(m, year)
	}

	//days of current month
	numDays += date - 1

	resultDay := (refDay + numDays) % 7
	return resultDay
}

func main() {
	startDate := []int{1, 1, 1901} //date,month,year
	endDate := []int{31, 12, 2000} //date,month,year

	dayOfWeek := calculateDayOfWeek(1900, 1, Monday, startDate[2], startDate[1], startDate[0])
	sundaysCnt := 0

	for y := startDate[2]; y <= endDate[2]; y++ {
		for m := 1; m <= 12; m++ {
			if dayOfWeek == Sunday {
				sundaysCnt++
			}

			daysInMonth := DaysInMonth(m, y)
			dayOfWeek = (daysInMonth + dayOfWeek) % 7
		}
	}

	fmt.Println("Total number of Sundays:", sundaysCnt)
}
