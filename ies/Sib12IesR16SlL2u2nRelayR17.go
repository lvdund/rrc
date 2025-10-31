package ies

import "rrc/utils"

// SIB12-IEs-r16-sl-L2U2N-Relay-r17 ::= ENUMERATED
type Sib12IesR16SlL2u2nRelayR17 struct {
	Value utils.ENUMERATED
}

const (
	Sib12IesR16SlL2u2nRelayR17EnumeratedNothing = iota
	Sib12IesR16SlL2u2nRelayR17EnumeratedEnabled
)
