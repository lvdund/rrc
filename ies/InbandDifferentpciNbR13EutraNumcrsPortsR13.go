package ies

import "rrc/utils"

// Inband-DifferentPCI-NB-r13-eutra-NumCRS-Ports-r13 ::= ENUMERATED
type InbandDifferentpciNbR13EutraNumcrsPortsR13 struct {
	Value utils.ENUMERATED
}

const (
	InbandDifferentpciNbR13EutraNumcrsPortsR13EnumeratedNothing = iota
	InbandDifferentpciNbR13EutraNumcrsPortsR13EnumeratedSame
	InbandDifferentpciNbR13EutraNumcrsPortsR13EnumeratedFour
)
