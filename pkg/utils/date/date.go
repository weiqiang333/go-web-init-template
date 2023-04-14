package date

import (
	"fmt"
	"time"
)

// GetNowDay 返回当前时间的日期 固定格式 2006-01-02
func GetNowDay() string {
	return time.Now().UTC().Format("2006-01-02")
}

// GetNowDate 返回当前时间的日期 固定格式 2006-01-02 15:04:05
func GetNowDate() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}

// GetDateYearMonth 返回日期年月 固定格式 2006-01-02 -> 2006-01
func GetDateYearMonth(date string) string {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, date)
	return t.Format("2006-01")
}

// GetDateYear 返回日期年 固定格式 2006-01-02 -> 2006
func GetDateYear(date string) string {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, date)
	return t.Format("2006")
}

// GetNowBeforeDate 获取当前时间偏移 固定格式 2006-01-02 15:04:05 (duration 支持正负)
func GetNowBeforeDate(duration time.Duration) string {
	return time.Now().UTC().Add(duration).Format("2006-01-02 15:04:05")
}

// GetBeforeDay 返回前后多少天的时间(7/-7) 固定格式 2006-01-02
func GetBeforeDay(i int) string {
	return time.Now().AddDate(0, 0, i).Format("2006-01-02")
}

// GetBeforeMonth 返回前后多少月的时间(3/-3) 固定格式 2006-01-02
func GetBeforeMonth(i int) string {
	LastMonth := time.Now().AddDate(0, i, 0)
	return LastMonth.Format("2006-01-02")
}

// GetLastMonth1stDay 获取上个月的1号时间 固定格式 2006-01-02
func GetLastMonth1stDay() string {
	LastMonth := time.Now().AddDate(0, 0, -30).Format("2006-01")
	layout := "2006-01"
	t, _ := time.Parse(layout, LastMonth)
	return t.Format("2006-01-02")
}

// GetDateRange 获取日期范围日期 list. req: ("2022-10-02", "2022-12-02")
func GetDateRange(dateStart, dateEnd string) []string {
	var dateList []string
	layout := "2006-01-02"
	start, _ := time.Parse(layout, dateStart)
	end, _ := time.Parse(layout, dateEnd)
	for d := start; d.After(end) == false; d = d.AddDate(0, 0, 1) {
		d.Format("2006-01-02")
		dateList = append(dateList, d.Format("2006-01-02"))
	}
	return dateList
}

// GetDateRangeYearToMonth 获取时间期间的所有月份 固定格式 2006-01-02
func GetDateRangeYearToMonth(dateStart, dateEnd string) ([]string, error) {
	var dateRangeMonth []string
	layout := "2006-01-02"
	t0, err := time.Parse(layout, dateStart)
	if err != nil {
		return dateRangeMonth, err
	}
	t1, err := time.Parse(layout, dateEnd)
	if err != nil {
		return dateRangeMonth, err
	}

	dateRangeMonth = append(dateRangeMonth, dateStart)
	for d := t0.AddDate(0, 1, 0); d.After(t1) == false; d = d.AddDate(0, 1, 0) {
		dateRangeMonth = append(dateRangeMonth, fmt.Sprintf("%s%s", d.Format("2006-01"), "-01"))
	}
	dateRangeMonth = append(dateRangeMonth, dateEnd)
	return dateRangeMonth, nil
}

// GetDateSubPeriodUnit 两个日期之间相差单位. "2022-10-01", "2023-10-1" -> year
func GetDateSubPeriodUnit(dateStart, dateEnd string) (string, error) {
	layout := "2006-01-02"
	t0, err := time.Parse(layout, dateStart)
	if err != nil {
		return "", err
	}
	t1, err := time.Parse(layout, dateEnd)
	if err != nil {
		return "", err
	}
	dateDiffer := t1.Sub(t0)
	if dateDiffer > time.Hour*24*360 {
		return "year", nil
	}
	if dateDiffer >= time.Hour*24*28 {
		return "month", nil
	}
	return "", fmt.Errorf("unrecognized time period")
}

// GetDateSubPeriodDays 两个时间之间相差天数.
func GetDateSubPeriodDays(dateStart, dateEnd string) (float64, error) {
	var day float64
	layout := "2006-01-02"
	t0, err := time.Parse(layout, dateStart)
	if err != nil {
		return day, err
	}
	t1, err := time.Parse(layout, dateEnd)
	if err != nil {
		return day, err
	}
	day = t1.Sub(t0).Hours() / 24
	return day, nil
}

// IfDateSize 判断时间大小，start >= end false, start < end true
func IfDateSize(start, end string) bool {
	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, start)
	endDate, _ := time.Parse(layout, end)
	return startDate.Before(endDate)
}
