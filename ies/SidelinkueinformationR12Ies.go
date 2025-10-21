package ies

import "rrc/utils"

// SidelinkUEInformation-r12-IEs ::= SEQUENCE
type SidelinkueinformationR12Ies struct {
	CommrxinterestedfreqR12  *ArfcnValueeutraR9
	CommtxresourcereqR12     *SlCommtxresourcereqR12
	DiscrxinterestR12        *utils.ENUMERATED
	DisctxresourcereqR12     *utils.INTEGER
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *SidelinkueinformationV1310Ies
}
