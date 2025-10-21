package ies

import "rrc/utils"

// RRCConnectionResumeComplete-v1610-IEs ::= SEQUENCE
type RrcconnectionresumecompleteV1610Ies struct {
	MeasresultlistidleR16    *MeasresultlistidleR15
	MeasresultlistextidleR16 *MeasresultlistextidleR16
	MeasresultlistidlenrR16  *MeasresultlistidlenrR16
	ScgConfigresponsenrR16   *utils.OCTETSTRING
	Noncriticalextension     *interface{}
}
