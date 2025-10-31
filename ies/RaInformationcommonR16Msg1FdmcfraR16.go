package ies

import "rrc/utils"

// RA-InformationCommon-r16-msg1-FDMCFRA-r16 ::= ENUMERATED
type RaInformationcommonR16Msg1FdmcfraR16 struct {
	Value utils.ENUMERATED
}

const (
	RaInformationcommonR16Msg1FdmcfraR16EnumeratedNothing = iota
	RaInformationcommonR16Msg1FdmcfraR16EnumeratedOne
	RaInformationcommonR16Msg1FdmcfraR16EnumeratedTwo
	RaInformationcommonR16Msg1FdmcfraR16EnumeratedFour
	RaInformationcommonR16Msg1FdmcfraR16EnumeratedEight
)
