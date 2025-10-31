package ies

import "rrc/utils"

// SRS-Resource-nrofSRS-Ports ::= ENUMERATED
type SrsResourceNrofsrsPorts struct {
	Value utils.ENUMERATED
}

const (
	SrsResourceNrofsrsPortsEnumeratedNothing = iota
	SrsResourceNrofsrsPortsEnumeratedPort1
	SrsResourceNrofsrsPortsEnumeratedPorts2
	SrsResourceNrofsrsPortsEnumeratedPorts4
)
