package ies

import "rrc/utils"

// RRCConnectionRelease-v1530-IEs ::= SEQUENCE
type RrcconnectionreleaseV1530Ies struct {
	DrbContinuerohcR15      *utils.ENUMERATED
	NexthopchainingcountR15 *Nexthopchainingcount
	MeasidleconfigR15       *MeasidleconfigdedicatedR15
	RrcInactiveconfigR15    *RrcInactiveconfigR15
	CnTypeR15               *utils.ENUMERATED
	Noncriticalextension    *RrcconnectionreleaseV1540Ies
}
