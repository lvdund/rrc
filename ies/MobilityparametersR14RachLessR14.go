package ies

import "rrc/utils"

// MobilityParameters-r14-rach-Less-r14 ::= ENUMERATED
type MobilityparametersR14RachLessR14 struct {
	Value utils.ENUMERATED
}

const (
	MobilityparametersR14RachLessR14EnumeratedNothing = iota
	MobilityparametersR14RachLessR14EnumeratedSupported
)
