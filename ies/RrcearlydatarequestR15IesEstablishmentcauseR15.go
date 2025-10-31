package ies

import "rrc/utils"

// RRCEarlyDataRequest-r15-IEs-establishmentCause-r15 ::= ENUMERATED
type RrcearlydatarequestR15IesEstablishmentcauseR15 struct {
	Value utils.ENUMERATED
}

const (
	RrcearlydatarequestR15IesEstablishmentcauseR15EnumeratedNothing = iota
	RrcearlydatarequestR15IesEstablishmentcauseR15EnumeratedMo_Data
	RrcearlydatarequestR15IesEstablishmentcauseR15EnumeratedDelaytolerantaccess
)
