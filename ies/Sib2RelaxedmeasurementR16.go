package ies

import "rrc/utils"

// SIB2-relaxedMeasurement-r16 ::= SEQUENCE
type Sib2RelaxedmeasurementR16 struct {
	LowmobilityevaluationR16       *Sib2RelaxedmeasurementR16LowmobilityevaluationR16
	CelledgeevaluationR16          *Sib2RelaxedmeasurementR16CelledgeevaluationR16
	CombinerelaxedmeasconditionR16 *utils.BOOLEAN
	HighprioritymeasrelaxR16       *utils.BOOLEAN
}
