package ies

import "rrc/utils"

// RACH-ConfigGeneric-ra-ResponseWindow-v1610 ::= ENUMERATED
type RachConfiggenericRaResponsewindowV1610 struct {
	Value utils.ENUMERATED
}

const (
	RachConfiggenericRaResponsewindowV1610EnumeratedNothing = iota
	RachConfiggenericRaResponsewindowV1610EnumeratedSl60
	RachConfiggenericRaResponsewindowV1610EnumeratedSl160
)
