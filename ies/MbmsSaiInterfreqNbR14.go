package ies

// MBMS-SAI-InterFreq-NB-r14 ::= SEQUENCE
type MbmsSaiInterfreqNbR14 struct {
	DlCarrierfreqR14     CarrierfreqNbR13
	MbmsSaiListR14       MbmsSaiListR11
	MultibandinfolistR14 *AdditionalbandinfolistNbR14
}
