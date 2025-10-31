package ies

import "rrc/utils"

// MBSFN-SubframeConfig ::= SEQUENCE
type MbsfnSubframeconfig struct {
	Radioframeallocationperiod MbsfnSubframeconfigRadioframeallocationperiod
	Radioframeallocationoffset utils.INTEGER `lb:0,ub:7`
	Subframeallocation         MbsfnSubframeconfigSubframeallocation
}
