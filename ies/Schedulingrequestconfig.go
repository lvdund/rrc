package ies

import "rrc/utils"

// SchedulingRequestConfig ::= CHOICE
type Schedulingrequestconfig interface {
	isSchedulingrequestconfig()
}

type SchedulingrequestconfigRelease struct {
	Value struct{}
}

func (*SchedulingrequestconfigRelease) isSchedulingrequestconfig() {}

type SchedulingrequestconfigSetup struct {
	Value interface{}
}

func (*SchedulingrequestconfigSetup) isSchedulingrequestconfig() {}
