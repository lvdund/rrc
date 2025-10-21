package ies

import "rrc/utils"

// UAC-AC1-SelectAssistInfo-r15 ::= ENUMERATED
type UacAc1SelectassistinfoR15 struct {
	Value utils.ENUMERATED
}

const (
	UacAc1SelectassistinfoR15A = 0
	UacAc1SelectassistinfoR15B = 1
	UacAc1SelectassistinfoR15C = 2
)
