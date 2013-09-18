package main

import (
	"testing"
)

func TestIsLeapYear(t *testing.T) {
	if IsLeapYear(2013) {
		t.Error("2013 is not a leap-year!")
	}
	if IsLeapYear(2100) {
		t.Error("2100 is not a leap-year!")
	}
	if !IsLeapYear(2012) {
		t.Error("2012 is a leap-year!")
	}
	if !IsLeapYear(2000) {
		t.Error("2000 is a leap-year!")
	}
}

func TestDayOfMonth(t *testing.T) {
	if DayOfMonth(2013, 1) != 31 {
		t.Error("Expected: 31, Actual: ", DayOfMonth(2013, 1))
	}
	if DayOfMonth(2013, 2) != 28 {
		t.Error("Expected: 28, Actual: ", DayOfMonth(2013, 2))
	}
	if DayOfMonth(2013, 4) != 30 {
		t.Error("Expected: 30, Actual: ", DayOfMonth(2013, 4))
	}

	if DayOfMonth(2012, 2) != 29 {
		t.Error("Expected: 29, Actual: ", DayOfMonth(2012, 2))
	}
}

func TestDayOfYear(t *testing.T) {
	if DayOfYear(2013, 1, 1) != 1 {
		t.Error("Expected: 1, Actual: ", DayOfYear(2013, 1, 1))
	}
	if DayOfYear(2013, 1, 31) != 31 {
		t.Error("Expected: 31, Actual: ", DayOfYear(2013, 1, 31))
	}
	if DayOfYear(2013, 2, 1) != 32 {
		t.Error("Expected: 32, Actual: ", DayOfYear(2013, 2, 1))
	}
	if DayOfYear(2013, 2, 28) != 59 {
		t.Error("Expected: 59, Actual: ", DayOfYear(2013, 2, 28))
	}
	if DayOfYear(2013, 3, 1) != 60 {
		t.Error("Expected: 60, Actual: ", DayOfYear(2013, 3, 1))
	}
	if DayOfYear(2013, 12, 31) != 365 {
		t.Error("Expected: 365, Actual: ", DayOfYear(2013, 12, 31))
	}

	if DayOfYear(2012, 2, 28) != 59 {
		t.Error("Expected: 59, Actual: ", DayOfYear(2012, 2, 28))
	}
	if DayOfYear(2012, 2, 29) != 60 {
		t.Error("Expected: 60, Actual: ", DayOfYear(2012, 2, 29))
	}
	if DayOfYear(2012, 3, 1) != 61 {
		t.Error("Expected: 61, Actual: ", DayOfYear(2012, 3, 1))
	}
	if DayOfYear(2012, 12, 31) != 366 {
		t.Error("Expected: 366, Actual: ", DayOfYear(2012, 12, 31))
	}
	//TODO 要エラー条件
}
