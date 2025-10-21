package ies

import "rrc/utils"

// SCG-ConfigInfo-v1530-IEs ::= SEQUENCE
type ScgConfiginfoV1530Ies struct {
	DrbToaddmodlistscgR15  *DrbInfolistscgR15
	DrbToreleaselistscgR15 *DrbToreleaselistR15
	Noncriticalextension   *interface{}
}
