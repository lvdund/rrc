package ies

import "rrc/utils"

// CG-ConfigInfo-v1610-IEs-dummy1 ::= SEQUENCE
type CgConfiginfoV1610IesDummy1 struct {
	FailuretypeeutraR16   CgConfiginfoV1610IesDummy1FailuretypeeutraR16
	MeasresultscgEutraR16 utils.OCTETSTRING
}
