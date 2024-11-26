package time_units


import (
"encoding/json"
"testing"
"time"
)

func TestParseTimeUnitDuration(t *testing.T) {
	tests := []struct {
		input    string
		expected TimeUnitDuration
		hasError bool
	}{
		{"2d3h45m30s", TimeUnitDuration{2, 3, 45, 30}, false},
		{"1d", TimeUnitDuration{1, 0, 0, 0}, false},
		{"5h", TimeUnitDuration{0, 5, 0, 0}, false},
		{"10m", TimeUnitDuration{0, 0, 10, 0}, false},
		{"20s", TimeUnitDuration{0, 0, 0, 20}, false},
		{"", TimeUnitDuration{}, false},
		{"invalid", TimeUnitDuration{}, true},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result, err := ParseTimeUnitDuration(test.input)
			if test.hasError {
				if err == nil {
					t.Errorf("expected error for input '%s', but got none", test.input)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected error for input '%s': %v", test.input, err)
				}
				if result.GetUnitDays() != test.expected.GetUnitDays() ||
					result.GetUnitHours() != test.expected.GetUnitHours() ||
					result.GetUnitMinutes() != test.expected.GetUnitMinutes() ||
					result.GetUnitSeconds() != test.expected.GetUnitSeconds() {
					t.Errorf("expected %+v, got %+v", test.expected, result)
				}
			}
		})
	}
}

func TestTimeUnitDuration_Duration(t *testing.T) {
	tests := []struct {
		duration TimeUnitDuration
		expected time.Duration
	}{
		{TimeUnitDuration{1, 0, 0, 0}, 24 * time.Hour},
		{TimeUnitDuration{0, 1, 0, 0}, time.Hour},
		{TimeUnitDuration{0, 0, 1, 0}, time.Minute},
		{TimeUnitDuration{0, 0, 0, 1}, time.Second},
		{TimeUnitDuration{1, 2, 30, 15}, 95415 * time.Second},
	}

	for _, test := range tests {
		t.Run(test.duration.String(), func(t *testing.T) {
			result := test.duration.Duration()
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

func TestTimeUnitDuration_MarshalJSON(t *testing.T) {
	duration := TimeUnitDuration{2, 3, 45, 30}
	expected := `"2d3h45m30s"`

	result, err := json.Marshal(duration)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(result) != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestTimeUnitDuration_UnmarshalJSON(t *testing.T) {
	input := `"2d3h45m30s"`
	expected := TimeUnitDuration{2, 3, 45, 30}

	var result TimeUnitDuration
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.GetUnitDays() != expected.GetUnitDays() ||
		result.GetUnitHours() != expected.GetUnitHours() ||
		result.GetUnitMinutes() != expected.GetUnitMinutes() ||
		result.GetUnitSeconds() != expected.GetUnitSeconds() {
		t.Errorf("expected %+v, got %+v", expected, result)
	}
}

func TestTimeUnitDuration_String(t *testing.T) {
	tests := []struct {
		duration TimeUnitDuration
		expected string
	}{
		{TimeUnitDuration{2, 3, 45, 30}, "2d3h45m30s"},
		{TimeUnitDuration{0, 0, 0, 1}, "1s"},
		{TimeUnitDuration{0, 0, 0, 0}, "0s"},
		{TimeUnitDuration{1, 0, 0, 0}, "1d"},
		{TimeUnitDuration{0, 1, 0, 0}, "1h"},
		{TimeUnitDuration{0, 0, 1, 0}, "1m"},
	}

	for _, test := range tests {
		t.Run(test.duration.String(), func(t *testing.T) {
			result := test.duration.String()
			if result != test.expected {
				t.Errorf("expected %s, got %s", test.expected, result)
			}
		})
	}
}
