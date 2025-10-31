package ies

import "rrc/utils"

// SIB-GuardbandInbandDiffPCI-TDD-NB-r15-sib-EUTRA-NumCRS-Ports-r15 ::= ENUMERATED
type SibGuardbandinbanddiffpciTddNbR15SibEutraNumcrsPortsR15 struct {
	Value utils.ENUMERATED
}

const (
	SibGuardbandinbanddiffpciTddNbR15SibEutraNumcrsPortsR15EnumeratedNothing = iota
	SibGuardbandinbanddiffpciTddNbR15SibEutraNumcrsPortsR15EnumeratedSame
	SibGuardbandinbanddiffpciTddNbR15SibEutraNumcrsPortsR15EnumeratedFour
)
