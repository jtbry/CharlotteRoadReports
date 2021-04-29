package common

import "time"

// Parse an ISO8601 Local (Eastern Time) to time.Time
func ParseIso8601Local(str string) (time.Time, error) {
	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		return time.Time{}, err
	}

	t, err := time.ParseInLocation("2006-01-02T15:04:05", str, location)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
