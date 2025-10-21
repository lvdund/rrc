package ies

import "rrc/utils"

// SystemInformationBlockType18-r12 ::= SEQUENCE
// Extensible
type Systeminformationblocktype18R12 struct {
	CommconfigR12            *interface{}
	Latenoncriticalextension *utils.OCTETSTRING
}
