package ies

import "rrc/utils"

// RRCConnectionRelease-NB-v1530-IEs ::= SEQUENCE
type RrcconnectionreleaseNbV1530Ies struct {
	DrbContinuerohcR15      *utils.ENUMERATED
	NexthopchainingcountR15 *Nexthopchainingcount
	Noncriticalextension    *RrcconnectionreleaseNbV1550Ies
}
