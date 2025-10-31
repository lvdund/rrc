package ies

import "rrc/utils"

// MBSFN-AreaInfo-r16-timeSeparation-r16 ::= ENUMERATED
type MbsfnAreainfoR16TimeseparationR16 struct {
	Value utils.ENUMERATED
}

const (
	MbsfnAreainfoR16TimeseparationR16EnumeratedNothing = iota
	MbsfnAreainfoR16TimeseparationR16EnumeratedSl2
	MbsfnAreainfoR16TimeseparationR16EnumeratedSl4
)
