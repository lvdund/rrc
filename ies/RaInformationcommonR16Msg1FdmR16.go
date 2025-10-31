package ies

import "rrc/utils"

// RA-InformationCommon-r16-msg1-FDM-r16 ::= ENUMERATED
type RaInformationcommonR16Msg1FdmR16 struct {
	Value utils.ENUMERATED
}

const (
	RaInformationcommonR16Msg1FdmR16EnumeratedNothing = iota
	RaInformationcommonR16Msg1FdmR16EnumeratedOne
	RaInformationcommonR16Msg1FdmR16EnumeratedTwo
	RaInformationcommonR16Msg1FdmR16EnumeratedFour
	RaInformationcommonR16Msg1FdmR16EnumeratedEight
)
