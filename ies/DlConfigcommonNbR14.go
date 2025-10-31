package ies

// DL-ConfigCommon-NB-r14 ::= SEQUENCE
// Extensible
type DlConfigcommonNbR14 struct {
	DlCarrierconfigR14 DlCarrierconfigcommonNbR14
	PcchConfigR14      *PcchConfigNbR14
}
