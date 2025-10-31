package ies

// CRS-AssistanceInfo-r11 ::= SEQUENCE
// Extensible
type CrsAssistanceinfoR11 struct {
	PhyscellidR11              Physcellid
	AntennaportscountR11       CrsAssistanceinfoR11AntennaportscountR11
	MbsfnSubframeconfiglistR11 MbsfnSubframeconfiglist
}
