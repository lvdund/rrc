package ies

import "rrc/utils"

// SystemInformationBlockType1-NB-v1450 ::= SEQUENCE
type Systeminformationblocktype1NbV1450 struct {
	NrsCrsPoweroffsetV1450 *utils.ENUMERATED
	Noncriticalextension   *Systeminformationblocktype1NbV1530
}
