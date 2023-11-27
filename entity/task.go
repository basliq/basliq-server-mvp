package entity

import "time"

type Task struct {
	Id                uint
	Title             string
	Description       string
	Assignee          User
	Subtasks          []Task
	DateOpened        time.Time
	DateClosed        time.Time
	DateCreated       time.Time
	DateUpdated       time.Time
	EstimatedDuration time.Duration
	ActualDuration    time.Duration
}
