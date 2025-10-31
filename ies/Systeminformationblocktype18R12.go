package ies

import "rrc/utils"

// SystemInformationBlockType18-r12 ::= SEQUENCE
// Extensible
type Systeminformationblocktype18R12 struct {
	CommconfigR12            *Systeminformationblocktype18R12CommconfigR12
	Latenoncriticalextension *utils.OCTETSTRING
}
