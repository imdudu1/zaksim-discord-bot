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

func Today(loc string) (time.Time, time.Time, error) {
	location, err := time.LoadLocation(loc)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	now := time.Now().In(location)
	return Bod(now), Truncate(now).AddDate(0, 0, 1), err
}

func LocNow(loc string) time.Time {
	location, err := time.LoadLocation(loc)
	if err != nil {
		return time.Time{}
	}
	return time.Now().In(location)
}
