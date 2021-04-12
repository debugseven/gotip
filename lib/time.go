package lib

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// data set time
type Time struct {
	Hours, Minutes uint8
	Seconds        uint32
}

// create a new time from seconds
func FromSeconds(seconds uint32) (Time, error) {
	h := uint8(seconds / SecondsPerHour)
	m := uint8((seconds / 60) % 60)
	err := checkBoundries(h, m)

	if err != nil {
		// return error
		return Time{}, errors.New(fmt.Sprintf("boundries check failed (%v)", err))
	}

	return Time{uint8(h), uint8(m), seconds}, nil
}

// create time from string format xx:xx
func FromTime(time string) (Time, error) {
	parts := strings.Split(time, ":")

	if len(parts) != 2 {
		return Time{}, errors.New("wrong format. expected: [hh:mm]")
	}

	h, m, s, err := parseTimeEntries(parts[0], parts[1])

	if err != nil {
		return Time{}, err
	}

	return Time{h, m, s}, nil
}

// get the current time data struct
func Current() (Time, error) {
	current := time.Now()
	h, m, s, err := parseTimeEntries(current.Format("15"), current.Format("04"))

	if err != nil {
		return Time{}, err
	}

	return Time{h, m, s}, nil
}

// the time data in percent
func (t Time) Percent() float32 {
	return float32(t.Seconds) / float32(SecondsPerDay)
}

// get info to the time data struct
func (t Time) Info() string {
	var hours, minutes string

	// fill empty 0 when less than 10 for hours
	if t.Hours < 10 {
		hours = fmt.Sprintf("0%v", t.Hours)
	} else {
		hours = fmt.Sprintf("%v", t.Hours)
	}

	// fill empty 0 when less than 10 for minutes
	if t.Minutes < 10 {
		minutes = fmt.Sprintf("0%v", t.Minutes)
	} else {
		minutes = fmt.Sprintf("%v", t.Minutes)
	}

	return fmt.Sprintf("%v:%v is %.2f%%", hours, minutes, t.Percent()*100)
}
