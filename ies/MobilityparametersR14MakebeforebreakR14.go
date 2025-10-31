package ies

import "rrc/utils"

// MobilityParameters-r14-makeBeforeBreak-r14 ::= ENUMERATED
type MobilityparametersR14MakebeforebreakR14 struct {
	Value utils.ENUMERATED
}

const (
	MobilityparametersR14MakebeforebreakR14EnumeratedNothing = iota
	MobilityparametersR14MakebeforebreakR14EnumeratedSupported
)
