package ies

import "rrc/utils"

// UAC-AC1-SelectAssistInfo-r16 ::= ENUMERATED
type UacAc1SelectassistinfoR16 struct {
	Value utils.ENUMERATED
}

const (
	UacAc1SelectassistinfoR16A             = 0
	UacAc1SelectassistinfoR16B             = 1
	UacAc1SelectassistinfoR16C             = 2
	UacAc1SelectassistinfoR16Notconfigured = 3
)
