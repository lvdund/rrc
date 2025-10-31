package ies

// SCG-ConfigInfo-v1530-IEs ::= SEQUENCE
type ScgConfiginfoV1530 struct {
	DrbToaddmodlistscgR15  *DrbInfolistscgR15
	DrbToreleaselistscgR15 *DrbToreleaselistR15
	Noncriticalextension   *ScgConfiginfoV1530IesNoncriticalextension
}
