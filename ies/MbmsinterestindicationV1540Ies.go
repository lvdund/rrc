package ies

import "rrc/utils"

// MBMSInterestIndication-v1540-IEs ::= SEQUENCE
type MbmsinterestindicationV1540Ies struct {
	MbmsRomInfolistR15   *interface{}
	Noncriticalextension *MbmsinterestindicationV1610Ies
}
