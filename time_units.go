package time_units

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// TimeUnitDuration represents a duration with days, hours, minutes, and seconds.
type TimeUnitDuration struct {
	days    int
	hours   int
	minutes int
	seconds int
}

var durationPattern = regexp.MustCompile(`^(\d+d)?(\d+h)?(\d+m)?(\d+s)?$`)

// ParseTimeUnitDuration converts a string in the format "10d5h6m1s" into a TimeUnitDuration.
func ParseTimeUnitDuration(input string) (TimeUnitDuration, error) {
	dur := TimeUnitDuration{}
	if !durationPattern.MatchString(input) {
		return dur, errors.New("invalid duration format, expected 'NdNhNmNs'")
	}

	daysMatch := regexp.MustCompile(`(\d+)d`).FindStringSubmatch(input)
	hoursMatch := regexp.MustCompile(`(\d+)h`).FindStringSubmatch(input)
	minutesMatch := regexp.MustCompile(`(\d+)m`).FindStringSubmatch(input)
	secondsMatch := regexp.MustCompile(`(\d+)s`).FindStringSubmatch(input)

	if len(daysMatch) > 0 {
		days, err := strconv.Atoi(daysMatch[1])
		if err != nil {
			return dur, err
		}
		dur.days = days
	}
	if len(hoursMatch) > 0 {
		hours, err := strconv.Atoi(hoursMatch[1])
		if err != nil {
			return dur, err
		}
		dur.hours = hours
	}
	if len(minutesMatch) > 0 {
		minutes, err := strconv.Atoi(minutesMatch[1])
		if err != nil {
			return dur, err
		}
		dur.minutes = minutes
	}
	if len(secondsMatch) > 0 {
		seconds, err := strconv.Atoi(secondsMatch[1])
		if err != nil {
			return dur, err
		}
		dur.seconds = seconds
	}

	return dur, nil
}

// GetUnitDays returns the number of days in the duration.
func (t TimeUnitDuration) GetUnitDays() int {
	return t.days
}

// GetUnitHours returns the number of hours in the duration.
func (t TimeUnitDuration) GetUnitHours() int {
	return t.hours
}

// GetUnitMinutes returns the number of minutes in the duration.
func (t TimeUnitDuration) GetUnitMinutes() int {
	return t.minutes
}

// GetUnitSeconds returns the number of seconds in the duration.
func (t TimeUnitDuration) GetUnitSeconds() int {
	return t.seconds
}

// Duration converts TimeUnitDuration to time.Duration.
func (t TimeUnitDuration) Duration() time.Duration {
	totalSeconds := t.seconds +
		t.minutes*60 +
		t.hours*3600 +
		t.days*86400
	return time.Duration(totalSeconds) * time.Second
}

// MarshalJSON converts TimeUnitDuration into its JSON string representation, excluding zero values.
func (t TimeUnitDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

// String returns a string representation of the TimeUnitDuration.
func (t TimeUnitDuration) String() string {
	var strBuilder strings.Builder

	if t.days > 0 {
		strBuilder.WriteString(fmt.Sprintf("%dd", t.days))
	}
	if t.hours > 0 {
		strBuilder.WriteString(fmt.Sprintf("%dh", t.hours))
	}
	if t.minutes > 0 {
		strBuilder.WriteString(fmt.Sprintf("%dm", t.minutes))
	}
	if t.seconds > 0 || (t.days == 0 && t.hours == 0 && t.minutes == 0) { // Ensure at least seconds are included if everything else is zero
		strBuilder.WriteString(fmt.Sprintf("%ds", t.seconds))
	}

	if strBuilder.Len() == 0 {
		// In case all values are zero, return "0s"
		strBuilder.WriteString("0s")
	}

	return strBuilder.String()
}

// UnmarshalJSON parses a JSON string into TimeUnitDuration.
func (t *TimeUnitDuration) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	parsed, err := ParseTimeUnitDuration(str)
	if err != nil {
		return err
	}

	*t = parsed
	return nil
}
