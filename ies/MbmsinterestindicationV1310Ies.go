package ies

import "rrc/utils"

// MBMSInterestIndication-v1310-IEs ::= SEQUENCE
type MbmsinterestindicationV1310Ies struct {
	MbmsServicesR13      *MbmsServicelistR13
	Noncriticalextension *MbmsinterestindicationV1540Ies
}
