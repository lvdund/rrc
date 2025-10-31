package ies

import "rrc/utils"

// PDCP-Config-discardTimer ::= ENUMERATED
type PdcpConfigDiscardtimer struct {
	Value utils.ENUMERATED
}

const (
	PdcpConfigDiscardtimerEnumeratedNothing = iota
	PdcpConfigDiscardtimerEnumeratedMs50
	PdcpConfigDiscardtimerEnumeratedMs100
	PdcpConfigDiscardtimerEnumeratedMs150
	PdcpConfigDiscardtimerEnumeratedMs300
	PdcpConfigDiscardtimerEnumeratedMs500
	PdcpConfigDiscardtimerEnumeratedMs750
	PdcpConfigDiscardtimerEnumeratedMs1500
	PdcpConfigDiscardtimerEnumeratedInfinity
)
