package ies

// SidelinkUEInformation-v1310-IEs ::= SEQUENCE
type SidelinkueinformationV1310 struct {
	CommtxresourcerequcR13        *SlCommtxresourcereqR12
	CommtxresourceinforeqrelayR13 *SidelinkueinformationV1310IesCommtxresourceinforeqrelayR13
	DisctxresourcereqV1310        *SidelinkueinformationV1310IesDisctxresourcereqV1310
	DisctxresourcereqpsR13        *SlDisctxresourcereqR13
	DiscrxgapreqR13               *SlGaprequestR13
	DisctxgapreqR13               *SlGaprequestR13
	DiscsysinforeportfreqlistR13  *SlDiscsysinforeportfreqlistR13
	Noncriticalextension          *SidelinkueinformationV1430
}
