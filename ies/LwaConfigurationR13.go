package ies

import "rrc/utils"

// LWA-Configuration-r13 ::= CHOICE
type LwaConfigurationR13 interface {
	isLwaConfigurationR13()
}

type LwaConfigurationR13Release struct {
	Value struct{}
}

func (*LwaConfigurationR13Release) isLwaConfigurationR13() {}

type LwaConfigurationR13Setup struct {
	Value interface{}
}

func (*LwaConfigurationR13Setup) isLwaConfigurationR13() {}
