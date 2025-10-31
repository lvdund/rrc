package ies

// MIMOParam-r17 ::= SEQUENCE
type MimoparamR17 struct {
	AdditionalpciToaddmodlistR17       *[]SsbMtcAdditionalpciR17 `lb:1,ub:maxNrofAdditionalPCIR17`
	AdditionalpciToreleaselistR17      *[]AdditionalpciindexR17  `lb:1,ub:maxNrofAdditionalPCIR17`
	UnifiedtciStatetypeR17             *MimoparamR17UnifiedtciStatetypeR17
	UplinkPowercontroltoaddmodlistR17  *[]UplinkPowercontrolR17   `lb:1,ub:maxULTciR17`
	UplinkPowercontroltoreleaselistR17 *[]UplinkPowercontrolidR17 `lb:1,ub:maxULTciR17`
	SfnschemepdcchR17                  *MimoparamR17SfnschemepdcchR17
	SfnschemepdschR17                  *MimoparamR17SfnschemepdschR17
}
