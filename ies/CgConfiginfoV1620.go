package ies

import "rrc/utils"

// CG-ConfigInfo-v1620-IEs ::= SEQUENCE
type CgConfiginfoV1620 struct {
	UeassistanceinformationsourcescgR16 *utils.OCTETSTRING
	Noncriticalextension                *CgConfiginfoV1640
}
