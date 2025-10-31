package ies

// UEInformationRequest-v1530-IEs ::= SEQUENCE
type UeinformationrequestV1530 struct {
	IdlemodemeasurementreqR15 *bool
	FlightpathinforeqR15      *FlightpathinforeportconfigR15
	Noncriticalextension      *UeinformationrequestV1530IesNoncriticalextension
}
