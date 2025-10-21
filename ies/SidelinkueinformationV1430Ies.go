package ies

import "rrc/utils"

// SidelinkUEInformation-v1430-IEs ::= SEQUENCE
type SidelinkueinformationV1430Ies struct {
	V2xCommrxinterestedfreqlistR14 *SlV2xCommfreqlistR14
	P2xCommtxtypeR14               *utils.ENUMERATED
	V2xCommtxresourcereqR14        *SlV2xCommtxfreqlistR14
	Noncriticalextension           *SidelinkueinformationV1530Ies
}
