package ies

// SL-PSSCH-TxConfig-r16 ::= SEQUENCE
// Extensible
type SlPsschTxconfigR16 struct {
	SlTypetxsyncR16             *SlTypetxsyncR16
	SlThresueSpeedR16           SlPsschTxconfigR16SlThresueSpeedR16
	SlParametersabovethresR16   SlPsschTxparametersR16
	SlParametersbelowthresR16   SlPsschTxparametersR16
	SlParametersabovethresV1650 *SlMinmaxmcsListR16 `ext`
	SlParametersbelowthresV1650 *SlMinmaxmcsListR16 `ext`
}
