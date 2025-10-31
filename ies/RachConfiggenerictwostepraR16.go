package ies

import "rrc/utils"

// RACH-ConfigGenericTwoStepRA-r16 ::= SEQUENCE
// Extensible
type RachConfiggenerictwostepraR16 struct {
	MsgaPrachConfigurationindexR16     *utils.INTEGER `lb:0,ub:262`
	MsgaRoFdmR16                       *RachConfiggenerictwostepraR16MsgaRoFdmR16
	MsgaRoFrequencystartR16            *utils.INTEGER `lb:0,ub:maxNrofPhysicalResourceBlocks1`
	MsgaZerocorrelationzoneconfigR16   *utils.INTEGER `lb:0,ub:15`
	MsgaPreamblepowerrampingstepR16    *RachConfiggenerictwostepraR16MsgaPreamblepowerrampingstepR16
	MsgaPreamblereceivedtargetpowerR16 *utils.INTEGER `lb:0,ub:-60`
	MsgbResponsewindowR16              *RachConfiggenerictwostepraR16MsgbResponsewindowR16
	PreambletransmaxR16                *RachConfiggenerictwostepraR16PreambletransmaxR16
	MsgbResponsewindowV1700            *RachConfiggenerictwostepraR16MsgbResponsewindowV1700 `ext`
}
