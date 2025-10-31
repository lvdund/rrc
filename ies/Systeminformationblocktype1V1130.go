package ies

// SystemInformationBlockType1-v1130-IEs ::= SEQUENCE
type Systeminformationblocktype1V1130 struct {
	TddConfigV1130         *TddConfigV1130
	CellselectioninfoV1130 *CellselectioninfoV1130
	Noncriticalextension   *Systeminformationblocktype1V1250
}
