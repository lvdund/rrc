package ies

import "rrc/utils"

// MeasResultGERAN-cgi-Info ::= SEQUENCE
type MeasresultgeranCgiInfo struct {
	Cellglobalid    Cellglobalidgeran
	Routingareacode *utils.BITSTRING `lb:8,ub:8`
}
