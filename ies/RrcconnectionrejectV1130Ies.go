package ies

import "rrc/utils"

// RRCConnectionReject-v1130-IEs ::= SEQUENCE
type RrcconnectionrejectV1130Ies struct {
	DeprioritisationreqR11 *interface{}
	Noncriticalextension   *RrcconnectionrejectV1320Ies
}
