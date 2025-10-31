package ies

import "rrc/utils"

// SCPTMConfiguration-NB-v1610-multiTB-Gap-r16 ::= ENUMERATED
type ScptmconfigurationNbV1610MultitbGapR16 struct {
	Value utils.ENUMERATED
}

const (
	ScptmconfigurationNbV1610MultitbGapR16EnumeratedNothing = iota
	ScptmconfigurationNbV1610MultitbGapR16EnumeratedSf16
	ScptmconfigurationNbV1610MultitbGapR16EnumeratedSf32
	ScptmconfigurationNbV1610MultitbGapR16EnumeratedSf64
	ScptmconfigurationNbV1610MultitbGapR16EnumeratedSf128
)
