package ies

// SidelinkUEInformation-v1530-IEs ::= SEQUENCE
type SidelinkueinformationV1530 struct {
	ReliabilityinfolistslR15 *SlReliabilitylistR15
	Noncriticalextension     *SidelinkueinformationV1530IesNoncriticalextension
}
