package ies

import "rrc/utils"

// NAICS-AssistanceInfo-r12 ::= CHOICE
type NaicsAssistanceinfoR12 interface {
	isNaicsAssistanceinfoR12()
}

type NaicsAssistanceinfoR12Release struct {
	Value struct{}
}

func (*NaicsAssistanceinfoR12Release) isNaicsAssistanceinfoR12() {}

type NaicsAssistanceinfoR12Setup struct {
	Value interface{}
}

func (*NaicsAssistanceinfoR12Setup) isNaicsAssistanceinfoR12() {}
