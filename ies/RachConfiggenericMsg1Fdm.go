package ies

import "rrc/utils"

// RACH-ConfigGeneric-msg1-FDM ::= ENUMERATED
type RachConfiggenericMsg1Fdm struct {
	Value utils.ENUMERATED
}

const (
	RachConfiggenericMsg1FdmEnumeratedNothing = iota
	RachConfiggenericMsg1FdmEnumeratedOne
	RachConfiggenericMsg1FdmEnumeratedTwo
	RachConfiggenericMsg1FdmEnumeratedFour
	RachConfiggenericMsg1FdmEnumeratedEight
)
