package ies

import "rrc/utils"

// CG-Config-v1620-IEs ::= SEQUENCE
type CgConfigV1620 struct {
	UeassistanceinformationscgR16 *utils.OCTETSTRING
	Noncriticalextension          *CgConfigV1630
}
