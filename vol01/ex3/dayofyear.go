package main

// IsLeapYearは、yearがうるう年の場合true、平年の場合はfalseを返します
func IsLeapYear(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// DayOfMonthは、指定のyear, monthに対する指定月の日数を返します
func DayOfMonth(year, month int) int {
	DAY_OF_MONTH := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if month == 2 && IsLeapYear(year) {
		return 29
	}

	return DAY_OF_MONTH[month-1]
}

// DayOfYearは、指定年月日がその年の何日目かを返します
func DayOfYear(year, month, day int) int {
	var days int
	for i := 1; i < month; i++ {
		days += DayOfMonth(year, i)
	}
	return day + days
}
