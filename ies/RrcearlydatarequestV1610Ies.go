package ies

import "rrc/utils"

// RRCEarlyDataRequest-v1610-IEs ::= SEQUENCE
type RrcearlydatarequestV1610Ies struct {
	EstablishmentcauseV1610 utils.ENUMERATED
	Noncriticalextension    *interface{}
}
