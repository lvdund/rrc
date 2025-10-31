package ies

import "rrc/utils"

// RA-InformationCommon-r16-msgA-RO-FDMCFRA-r17 ::= ENUMERATED
type RaInformationcommonR16MsgaRoFdmcfraR17 struct {
	Value utils.ENUMERATED
}

const (
	RaInformationcommonR16MsgaRoFdmcfraR17EnumeratedNothing = iota
	RaInformationcommonR16MsgaRoFdmcfraR17EnumeratedOne
	RaInformationcommonR16MsgaRoFdmcfraR17EnumeratedTwo
	RaInformationcommonR16MsgaRoFdmcfraR17EnumeratedFour
	RaInformationcommonR16MsgaRoFdmcfraR17EnumeratedEight
)
