package ies

import "rrc/utils"

// RateMatchPatternLTE-CRS-nrofCRS-Ports ::= ENUMERATED
type RatematchpatternlteCrsNrofcrsPorts struct {
	Value utils.ENUMERATED
}

const (
	RatematchpatternlteCrsNrofcrsPortsEnumeratedNothing = iota
	RatematchpatternlteCrsNrofcrsPortsEnumeratedN1
	RatematchpatternlteCrsNrofcrsPortsEnumeratedN2
	RatematchpatternlteCrsNrofcrsPortsEnumeratedN4
)
