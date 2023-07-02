package models

type Status int

const (
	Initialized Status = iota
	InProgres
	Blocked
	Finished
	Reopen
)
