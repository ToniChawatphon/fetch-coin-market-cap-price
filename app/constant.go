package app

import (
	"time"
)

var (
	Setting Config
	Run     *Service
	Agent   *Scheduler

	Sunday    = time.Weekday(0).String()
	Monday    = time.Weekday(1).String()
	Tuesday   = time.Weekday(2).String()
	Wednesday = time.Weekday(3).String()
	Thursday  = time.Weekday(4).String()
	Friday    = time.Weekday(5).String()
	Saturday  = time.Weekday(6).String()
)

const (
	// Const
	DATETIME_FORMAT    = "2006/01/02"
	CURRENTTIME_FORMAT = "20060102150405"

	// Path
	OutJson    = "out/coins_market_value.json"
	ConfigFile = "config/config.yaml"

	// Scheduler
	Second = "seconds"
	Minute = "minutes"
	Hour   = "hours"
	Day    = "days"
	Week   = "weeks"
	Month  = "months"
)
