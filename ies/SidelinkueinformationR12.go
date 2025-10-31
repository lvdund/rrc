package ies

import "rrc/utils"

// SidelinkUEInformation-r12-IEs ::= SEQUENCE
type SidelinkueinformationR12 struct {
	CommrxinterestedfreqR12  *ArfcnValueeutraR9
	CommtxresourcereqR12     *SlCommtxresourcereqR12
	DiscrxinterestR12        *bool
	DisctxresourcereqR12     *utils.INTEGER `lb:0,ub:63`
	Latenoncriticalextension *utils.OCTETSTRING
	Noncriticalextension     *SidelinkueinformationV1310
}
