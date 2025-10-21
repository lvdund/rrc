package ies

import "rrc/utils"

// SidelinkUEInformation-v1310-IEs ::= SEQUENCE
type SidelinkueinformationV1310Ies struct {
	CommtxresourcerequcR13        *SlCommtxresourcereqR12
	CommtxresourceinforeqrelayR13 *interface{}
	DisctxresourcereqV1310        *interface{}
	DisctxresourcereqpsR13        *SlDisctxresourcereqR13
	DiscrxgapreqR13               *SlGaprequestR13
	DisctxgapreqR13               *SlGaprequestR13
	DiscsysinforeportfreqlistR13  *SlDiscsysinforeportfreqlistR13
	Noncriticalextension          *SidelinkueinformationV1430Ies
}
