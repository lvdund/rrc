package ies

import "rrc/utils"

// BandNR-condHandoverTwoTriggerEvents-r16 ::= ENUMERATED
type BandnrCondhandovertwotriggereventsR16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrCondhandovertwotriggereventsR16EnumeratedNothing = iota
	BandnrCondhandovertwotriggereventsR16EnumeratedSupported
)
