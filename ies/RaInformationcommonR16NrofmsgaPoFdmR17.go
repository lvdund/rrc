package ies

import "rrc/utils"

// RA-InformationCommon-r16-nrofMsgA-PO-FDM-r17 ::= ENUMERATED
type RaInformationcommonR16NrofmsgaPoFdmR17 struct {
	Value utils.ENUMERATED
}

const (
	RaInformationcommonR16NrofmsgaPoFdmR17EnumeratedNothing = iota
	RaInformationcommonR16NrofmsgaPoFdmR17EnumeratedOne
	RaInformationcommonR16NrofmsgaPoFdmR17EnumeratedTwo
	RaInformationcommonR16NrofmsgaPoFdmR17EnumeratedFour
	RaInformationcommonR16NrofmsgaPoFdmR17EnumeratedEight
)
