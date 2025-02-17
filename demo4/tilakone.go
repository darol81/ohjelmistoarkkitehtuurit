package main

import "fmt"

type DigitalClock struct {
	hour       int
	minute     int
	mode_index int // 0: hour, 1: minute
}

func MakeDigitalClock() *DigitalClock {
	dc := new(DigitalClock)
	dc.hour = 0
	dc.minute = 0
	return dc
}

func (dc *DigitalClock) Inc() {
	if dc.mode_index == 0 {
		dc.hour = (dc.hour + 1) % 24
	} else {
		dc.minute = (dc.minute + 1) % 60
	}
}

func (dc *DigitalClock) Set() {
	dc.mode_index = (dc.mode_index + 1) % 2
}

func (dc *DigitalClock) GetTimeString() string {
	return fmt.Sprintf("%02d:%02d", dc.hour, dc.minute)
}

func main() {
	dc := MakeDigitalClock()
	dc.Inc()
	dc.Inc()
	dc.Set()
	dc.Inc()
	dc.Inc()
	dc.Inc()
	dc.Set()
	if dc.GetTimeString() == "02:03" {
		fmt.Println("pass")
	} else {
		fmt.Println("fail")
	}
}
