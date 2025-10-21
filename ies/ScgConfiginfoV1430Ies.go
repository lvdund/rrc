package ies

import "rrc/utils"

// SCG-ConfigInfo-v1430-IEs ::= SEQUENCE
type ScgConfiginfoV1430Ies struct {
	MakebeforebreakscgReqR14 *utils.ENUMERATED
	MeasgapconfigperccList   *MeasgapconfigperccListR14
	Noncriticalextension     *ScgConfiginfoV1530Ies
}
