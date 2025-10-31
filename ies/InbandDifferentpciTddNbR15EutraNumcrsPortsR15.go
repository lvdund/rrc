package ies

import "rrc/utils"

// Inband-DifferentPCI-TDD-NB-r15-eutra-NumCRS-Ports-r15 ::= ENUMERATED
type InbandDifferentpciTddNbR15EutraNumcrsPortsR15 struct {
	Value utils.ENUMERATED
}

const (
	InbandDifferentpciTddNbR15EutraNumcrsPortsR15EnumeratedNothing = iota
	InbandDifferentpciTddNbR15EutraNumcrsPortsR15EnumeratedSame
	InbandDifferentpciTddNbR15EutraNumcrsPortsR15EnumeratedFour
)
