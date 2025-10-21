package ies

import "rrc/utils"

// Enable256QAM-r14 ::= CHOICE
type Enable256qamR14 interface {
	isEnable256qamR14()
}

type Enable256qamR14Release struct {
	Value struct{}
}

func (*Enable256qamR14Release) isEnable256qamR14() {}

type Enable256qamR14Setup struct {
	Value interface{}
}

func (*Enable256qamR14Setup) isEnable256qamR14() {}
