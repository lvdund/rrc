package ies

import "rrc/utils"

// SchedulingRequestConfig-v1530 ::= CHOICE
type SchedulingrequestconfigV1530 interface {
	isSchedulingrequestconfigV1530()
}

type SchedulingrequestconfigV1530Release struct {
	Value struct{}
}

func (*SchedulingrequestconfigV1530Release) isSchedulingrequestconfigV1530() {}

type SchedulingrequestconfigV1530Setup struct {
	Value interface{}
}

func (*SchedulingrequestconfigV1530Setup) isSchedulingrequestconfigV1530() {}
