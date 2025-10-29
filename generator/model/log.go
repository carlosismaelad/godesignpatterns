package model

import "time"

type Log struct {
	Level   string
	Message string
	Time    time.Time
}