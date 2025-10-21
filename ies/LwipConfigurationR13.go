package ies

import "rrc/utils"

// LWIP-Configuration-r13 ::= CHOICE
type LwipConfigurationR13 interface {
	isLwipConfigurationR13()
}

type LwipConfigurationR13Release struct {
	Value struct{}
}

func (*LwipConfigurationR13Release) isLwipConfigurationR13() {}

type LwipConfigurationR13Setup struct {
	Value interface{}
}

func (*LwipConfigurationR13Setup) isLwipConfigurationR13() {}
