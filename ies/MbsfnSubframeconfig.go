package ies

import "rrc/utils"

// MBSFN-SubframeConfig ::= SEQUENCE
type MbsfnSubframeconfig struct {
	Radioframeallocationperiod utils.ENUMERATED
	Radioframeallocationoffset utils.INTEGER
	Subframeallocation         interface{}
}
