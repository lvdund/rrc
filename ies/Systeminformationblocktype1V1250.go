package ies

// SystemInformationBlockType1-v1250-IEs ::= SEQUENCE
type Systeminformationblocktype1V1250 struct {
	CellaccessrelatedinfoV1250   *Systeminformationblocktype1V1250IesCellaccessrelatedinfoV1250
	CellselectioninfoV1250       *CellselectioninfoV1250
	FreqbandindicatorpriorityR12 *bool
	Noncriticalextension         *Systeminformationblocktype1V1310
}
