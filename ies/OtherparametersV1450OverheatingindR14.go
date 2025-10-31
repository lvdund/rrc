package ies

import "rrc/utils"

// OtherParameters-v1450-overheatingInd-r14 ::= ENUMERATED
type OtherparametersV1450OverheatingindR14 struct {
	Value utils.ENUMERATED
}

const (
	OtherparametersV1450OverheatingindR14EnumeratedNothing = iota
	OtherparametersV1450OverheatingindR14EnumeratedSupported
)
