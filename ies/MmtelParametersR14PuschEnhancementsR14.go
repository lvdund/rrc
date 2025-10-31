package ies

import "rrc/utils"

// MMTEL-Parameters-r14-pusch-Enhancements-r14 ::= ENUMERATED
type MmtelParametersR14PuschEnhancementsR14 struct {
	Value utils.ENUMERATED
}

const (
	MmtelParametersR14PuschEnhancementsR14EnumeratedNothing = iota
	MmtelParametersR14PuschEnhancementsR14EnumeratedSupported
)
