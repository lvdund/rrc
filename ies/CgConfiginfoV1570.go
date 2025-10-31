package ies

// CG-ConfigInfo-v1570-IEs ::= SEQUENCE
type CgConfiginfoV1570 struct {
	SftdfrequencylistNr    *SftdFrequencylistNr
	SftdfrequencylistEutra *SftdFrequencylistEutra
	Noncriticalextension   *CgConfiginfoV1590
}
