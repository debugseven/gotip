package lib

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	SecondsPerDay  = 60 * 60 * 24
	SecondsPerHour = 60 * 60
)

// check bounds from 00:00 to 23:59
func checkBoundries(hours uint8, minutes uint8) error {
	if hours < 0 || 23 < hours {
		return errors.New("hours out of bounds")
	} else if minutes < 0 || 59 < minutes {
		return errors.New("minutes out of bounds")
	}
	return nil
}

// returns hours, minutes, seconds
func parseTimeEntries(hours string, minutes string) (uint8, uint8, uint32, error) {
	// parse hours
	x, err := strconv.Atoi(hours)
	if err != nil {
		return 0, 0, 0, err
	}

	h := uint8(x)

	// parse minutes
	x, err = strconv.Atoi(minutes)
	if err != nil {
		return 0, 0, 0, err
	}

	m := uint8(x)

	if err != nil {
		return 0, 0, 0, errors.New(fmt.Sprintf("boundries check failed (%v)", err))
	}

	// calculate seconds
	var s uint32 = uint32(h)*SecondsPerHour + uint32(m)*60
	return h, m, s, nil
}
