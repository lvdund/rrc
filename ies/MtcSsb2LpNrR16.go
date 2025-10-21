package ies

import "rrc/utils"

// MTC-SSB2-LP-NR-r16 ::= SEQUENCE
type MtcSsb2LpNrR16 struct {
	PciListR16     *interface{}
	PeriodicityR16 utils.ENUMERATED
}
