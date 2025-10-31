package ies

import "rrc/utils"

// HighSpeedEnhParameters-r14-prach-Enhancements-r14 ::= ENUMERATED
type HighspeedenhparametersR14PrachEnhancementsR14 struct {
	Value utils.ENUMERATED
}

const (
	HighspeedenhparametersR14PrachEnhancementsR14EnumeratedNothing = iota
	HighspeedenhparametersR14PrachEnhancementsR14EnumeratedSupported
)
