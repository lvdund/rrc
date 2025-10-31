package ies

// SidelinkUEInformation-v1430-IEs ::= SEQUENCE
type SidelinkueinformationV1430 struct {
	V2xCommrxinterestedfreqlistR14 *SlV2xCommfreqlistR14
	P2xCommtxtypeR14               *bool
	V2xCommtxresourcereqR14        *SlV2xCommtxfreqlistR14
	Noncriticalextension           *SidelinkueinformationV1530
}
