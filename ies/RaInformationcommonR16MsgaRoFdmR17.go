package ies

import "rrc/utils"

// RA-InformationCommon-r16-msgA-RO-FDM-r17 ::= ENUMERATED
type RaInformationcommonR16MsgaRoFdmR17 struct {
	Value utils.ENUMERATED
}

const (
	RaInformationcommonR16MsgaRoFdmR17EnumeratedNothing = iota
	RaInformationcommonR16MsgaRoFdmR17EnumeratedOne
	RaInformationcommonR16MsgaRoFdmR17EnumeratedTwo
	RaInformationcommonR16MsgaRoFdmR17EnumeratedFour
	RaInformationcommonR16MsgaRoFdmR17EnumeratedEight
)
