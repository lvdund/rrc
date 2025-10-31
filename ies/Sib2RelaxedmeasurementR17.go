package ies

import "rrc/utils"

// SIB2-relaxedMeasurement-r17 ::= SEQUENCE
type Sib2RelaxedmeasurementR17 struct {
	StationarymobilityevaluationR17      Sib2RelaxedmeasurementR17StationarymobilityevaluationR17
	CelledgeevaluationwhilestationaryR17 *Sib2RelaxedmeasurementR17CelledgeevaluationwhilestationaryR17
	Combinerelaxedmeascondition2R17      *utils.BOOLEAN
}
