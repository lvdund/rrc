package ies

// SCG-ConfigPartSCG-r12 ::= SEQUENCE
// Extensible
type ScgConfigpartscgR12 struct {
	RadioresourceconfigdedicatedscgR12 *RadioresourceconfigdedicatedscgR12
	ScelltoreleaselistscgR12           *ScelltoreleaselistR10
	PscelltoaddmodR12                  *PscelltoaddmodR12
	ScelltoaddmodlistscgR12            *ScelltoaddmodlistR10
	MobilitycontrolinfoscgR12          *MobilitycontrolinfoscgR12
}
