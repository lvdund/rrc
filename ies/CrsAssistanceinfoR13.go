package ies

// CRS-AssistanceInfo-r13 ::= SEQUENCE
// Extensible
type CrsAssistanceinfoR13 struct {
	PhyscellidR13              Physcellid
	AntennaportscountR13       CrsAssistanceinfoR13AntennaportscountR13
	MbsfnSubframeconfiglistR13 *MbsfnSubframeconfiglist
}
