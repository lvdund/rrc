package ies

import "rrc/utils"

// UAC-AC1-SelectAssistInfo-r15 ::= ENUMERATED
type UacAc1SelectassistinfoR15 struct {
	Value utils.ENUMERATED
}

const (
	UacAc1SelectassistinfoR15EnumeratedNothing = iota
	UacAc1SelectassistinfoR15EnumeratedA
	UacAc1SelectassistinfoR15EnumeratedB
	UacAc1SelectassistinfoR15EnumeratedC
)
