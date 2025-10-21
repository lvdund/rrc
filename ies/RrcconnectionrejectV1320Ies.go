package ies

import "rrc/utils"

// RRCConnectionReject-v1320-IEs ::= SEQUENCE
type RrcconnectionrejectV1320Ies struct {
	RrcSuspendindicationR13 *utils.ENUMERATED
	Noncriticalextension    *interface{}
}
