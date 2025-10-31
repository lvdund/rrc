package ies

// SystemInformationBlockType1-v1430-IEs ::= SEQUENCE
type Systeminformationblocktype1V1430 struct {
	EcalloverimsSupportR14       *bool
	TddConfigV1430               *TddConfigV1430
	CellaccessrelatedinfolistR14 *[]CellaccessrelatedinfoR14 `lb:1,ub:maxPLMN1R14`
	Noncriticalextension         *Systeminformationblocktype1V1450
}
