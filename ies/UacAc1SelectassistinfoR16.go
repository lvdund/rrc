package ies

import "rrc/utils"

// UAC-AC1-SelectAssistInfo-r16 ::= ENUMERATED
type UacAc1SelectassistinfoR16 struct {
	Value utils.ENUMERATED
}

const (
	UacAc1SelectassistinfoR16EnumeratedNothing = iota
	UacAc1SelectassistinfoR16EnumeratedA
	UacAc1SelectassistinfoR16EnumeratedB
	UacAc1SelectassistinfoR16EnumeratedC
	UacAc1SelectassistinfoR16EnumeratedNotconfigured
)
