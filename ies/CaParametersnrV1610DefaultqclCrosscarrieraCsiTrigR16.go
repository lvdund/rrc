package ies

import "rrc/utils"

// CA-ParametersNR-v1610-defaultQCL-CrossCarrierA-CSI-Trig-r16 ::= ENUMERATED
type CaParametersnrV1610DefaultqclCrosscarrieraCsiTrigR16 struct {
	Value utils.ENUMERATED
}

const (
	CaParametersnrV1610DefaultqclCrosscarrieraCsiTrigR16EnumeratedNothing = iota
	CaParametersnrV1610DefaultqclCrosscarrieraCsiTrigR16EnumeratedDiffonly
	CaParametersnrV1610DefaultqclCrosscarrieraCsiTrigR16EnumeratedBoth
)
