//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minutes, seconds int
}

type TimeParseError struct {
	msg   string
	input string
}

func (e *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", e.msg, e.input)
}

func ParseTime(input string) (Time, error) {
	components := strings.Split(input, ":")

	if len(components) != 3 {
		return Time{}, &TimeParseError{"invalid number of time components", input}
	} else {
		hour, err := strconv.Atoi(components[0])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing hour: %v", err), input}
		}
		minutes, err := strconv.Atoi(components[1])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing minutes: %v", err), input}
		}
		seconds, err := strconv.Atoi(components[2])
		if err != nil {
			return Time{}, &TimeParseError{fmt.Sprintf("Error parsing seconds: %v", err), input}
		}

		if hour > 23 || hour < 0 {
			return Time{}, &TimeParseError{"hour out of range: 0 - 23", input}
		}

		if minutes > 59 || minutes < 0 {
			return Time{}, &TimeParseError{"minutes out of range: 0 - 59", input}
		}

		if seconds > 59 || seconds < 0 {
			return Time{}, &TimeParseError{"seconds out of range: 0 - 59", input}
		}

		return Time{hour, minutes, seconds}, nil
	}
}
