package ies

import "rrc/utils"

// SchedulingRequestConfigSCell-r13 ::= CHOICE
type SchedulingrequestconfigscellR13 interface {
	isSchedulingrequestconfigscellR13()
}

type SchedulingrequestconfigscellR13Release struct {
	Value struct{}
}

func (*SchedulingrequestconfigscellR13Release) isSchedulingrequestconfigscellR13() {}

type SchedulingrequestconfigscellR13Setup struct {
	Value interface{}
}

func (*SchedulingrequestconfigscellR13Setup) isSchedulingrequestconfigscellR13() {}
