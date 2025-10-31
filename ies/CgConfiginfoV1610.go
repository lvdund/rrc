package ies

import "rrc/utils"

// CG-ConfigInfo-v1610-IEs ::= SEQUENCE
type CgConfiginfoV1610 struct {
	DrxInfomcg2                   *DrxInfo2
	AligneddrxIndication          *utils.BOOLEAN
	ScgfailureinfoR16             *CgConfiginfoV1610IesScgfailureinfoR16
	Dummy1                        *CgConfiginfoV1610IesDummy1
	SidelinkueinformationnrR16    *utils.OCTETSTRING
	SidelinkueinformationeutraR16 *utils.OCTETSTRING
	Noncriticalextension          *CgConfiginfoV1620
}
