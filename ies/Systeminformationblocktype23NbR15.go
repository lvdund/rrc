package ies

import "rrc/utils"

// SystemInformationBlockType23-NB-r15 ::= SEQUENCE
// Extensible
type Systeminformationblocktype23NbR15 struct {
	UlConfiglistV1530        *UlConfigcommonlistNbV1530
	UlConfiglistmixedV1530   *UlConfigcommonlistNbV1530
	Latenoncriticalextension *utils.OCTETSTRING
}
