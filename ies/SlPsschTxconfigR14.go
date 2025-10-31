package ies

// SL-PSSCH-TxConfig-r14 ::= SEQUENCE
// Extensible
type SlPsschTxconfigR14 struct {
	TypetxsyncR14           *SlTypetxsyncR14
	ThresueSpeedR14         SlPsschTxconfigR14ThresueSpeedR14
	ParametersabovethresR14 SlPsschTxparametersR14
	ParametersbelowthresR14 SlPsschTxparametersR14
}
