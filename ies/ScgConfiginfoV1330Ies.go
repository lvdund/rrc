package ies

import "rrc/utils"

// SCG-ConfigInfo-v1330-IEs ::= SEQUENCE
type ScgConfiginfoV1330Ies struct {
	MeasresultlistrssiScgR13 *MeasresultlistrssiScgR13
	Noncriticalextension     *ScgConfiginfoV1430Ies
}
