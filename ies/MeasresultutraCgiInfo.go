package ies

import "rrc/utils"

// MeasResultUTRA-cgi-Info ::= SEQUENCE
type MeasresultutraCgiInfo struct {
	Cellglobalid     Cellglobalidutra
	Locationareacode *utils.BITSTRING `lb:16,ub:16`
	Routingareacode  *utils.BITSTRING `lb:8,ub:8`
	PlmnIdentitylist *PlmnIdentitylist2
}
