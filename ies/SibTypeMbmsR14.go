package ies

import "rrc/utils"

// SIB-Type-MBMS-r14 ::= utils.ENUMERATED // Extensible
type SibTypeMbmsR14 struct {
	Value utils.ENUMERATED
}

const (
	SibTypeMbmsR14EnumeratedNothing = iota
	SibTypeMbmsR14EnumeratedSibtype10
	SibTypeMbmsR14EnumeratedSibtype11
	SibTypeMbmsR14EnumeratedSibtype12_V920
	SibTypeMbmsR14EnumeratedSibtype13_V920
	SibTypeMbmsR14EnumeratedSibtype15_V1130
	SibTypeMbmsR14EnumeratedSibtype16_V1130
)
