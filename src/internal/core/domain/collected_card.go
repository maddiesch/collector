package domain

import "time"

type CollectedCard struct {
	GroupName       string
	Name            string
	SetName         string
	CollectorNumber string
	IsFoil          bool
	Language        string
	Condition       CardCondition
	CreatedAt       time.Time
}
