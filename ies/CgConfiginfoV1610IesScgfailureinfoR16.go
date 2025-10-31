package ies

import "rrc/utils"

// CG-ConfigInfo-v1610-IEs-scgFailureInfo-r16 ::= SEQUENCE
type CgConfiginfoV1610IesScgfailureinfoR16 struct {
	FailuretypeR16   CgConfiginfoV1610IesScgfailureinfoR16FailuretypeR16
	MeasresultscgR16 utils.OCTETSTRING
}
