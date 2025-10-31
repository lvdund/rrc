package ies

import "rrc/utils"

// BandNR-condPSCellChangeTwoTriggerEvents-r16 ::= ENUMERATED
type BandnrCondpscellchangetwotriggereventsR16 struct {
	Value utils.ENUMERATED
}

const (
	BandnrCondpscellchangetwotriggereventsR16EnumeratedNothing = iota
	BandnrCondpscellchangetwotriggereventsR16EnumeratedSupported
)
