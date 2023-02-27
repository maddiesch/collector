package domain

import "time"

type CollectedCard struct {
	ID              string
	GroupName       string
	Name            string
	SetName         string
	CollectorNumber string
	IsFoil          bool
	Language        string
	Condition       CardCondition
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
