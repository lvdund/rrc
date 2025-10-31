package ies

// SystemInformationBlockType1-v920-IEs ::= SEQUENCE
type Systeminformationblocktype1V920 struct {
	ImsEmergencysupportR9 *bool
	CellselectioninfoV920 *CellselectioninfoV920
	Noncriticalextension  *Systeminformationblocktype1V1130
}
