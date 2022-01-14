package main

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	tt := []struct{
		name string
		dateString string
		time time.Time
	}{
		{name:"< day < month", dateString: "2001-03-01", time: time.Date(2001, 3, 1, 0, 0, 0, 0, time.UTC)},
		{name:"< day > month", dateString: "2001-03-10", time: time.Date(2001, 3, 10, 0, 0, 0, 0, time.UTC)},
		{name:"> day < month", dateString: "2001-10-02", time: time.Date(2001, 10, 2, 0, 0, 0, 0, time.UTC)},
		{name:"> day > month", dateString: "2001-11-12", time: time.Date(2001, 11, 12, 0, 0, 0, 0, time.UTC)},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ti := parseTime(tc.dateString)
			if ti != tc.time {
				t.Errorf("expected %v got %v", tc.time, ti)
			}
		})
	}
}

func TestParseStringTime(t *testing.T) {
	tt := []struct{
		name string
		year int
		month int
		day int
		out string
	}{
		{name: "YYYY-M-DD", year: 2001, month: 3, day: 12, out: "2001-03-12"},
		{name: "YYYY-MM-D", year: 2001, month: 10, day: 2, out: "2001-10-02"},
		{name: "YYYY-MM-D", year: 2001, month: 1, day: 2, out: "2001-01-02"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			m := time.Month(tc.month)
			o := parseStringTime(tc.year, m, tc.day)
			if o != tc.out {
				t.Errorf("expected '%v' got '%v'", tc.out, o)
			}
		})
	}
}
