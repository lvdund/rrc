package ies

import "rrc/utils"

// EUTRA-MBSFN-SubframeConfig ::= SEQUENCE
// Extensible
type EutraMbsfnSubframeconfig struct {
	Radioframeallocationperiod EutraMbsfnSubframeconfigRadioframeallocationperiod
	Radioframeallocationoffset utils.INTEGER `lb:0,ub:7`
	Subframeallocation1        EutraMbsfnSubframeconfigSubframeallocation1
	Subframeallocation2        *EutraMbsfnSubframeconfigSubframeallocation2
}
