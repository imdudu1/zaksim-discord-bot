package tiktok

import "time"

// ref: https://stackoverflow.com/questions/25254443/return-local-beginning-of-day-time-object

func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func Truncate(t time.Time) time.Time {
	return t.Truncate(24 * time.Hour)
}

// Today 오늘 하루를 나타내는 시간의 범위를 반환하는 함수
func Today(loc string) (time.Time, time.Time, error) {
	location, err := time.LoadLocation(loc)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	now := time.Now().In(location)
	return Bod(now), Truncate(now).AddDate(0, 0, 1), err
}

// LocaleNow 타임존 문자열을 이용해 Location 값을 만드는 함수
func LocaleNow(loc string) (time.Time, error) {
	location, err := time.LoadLocation(loc)
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().In(location), nil
}
