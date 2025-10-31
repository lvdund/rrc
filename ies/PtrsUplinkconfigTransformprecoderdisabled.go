package ies

import "rrc/utils"

// PTRS-UplinkConfig-transformPrecoderDisabled ::= SEQUENCE
type PtrsUplinkconfigTransformprecoderdisabled struct {
	Frequencydensity      *[]utils.INTEGER `lb:2,ub:2`
	Timedensity           *[]utils.INTEGER `lb:3,ub:3`
	Maxnrofports          PtrsUplinkconfigTransformprecoderdisabledMaxnrofports
	Resourceelementoffset *PtrsUplinkconfigTransformprecoderdisabledResourceelementoffset
	PtrsPower             PtrsUplinkconfigTransformprecoderdisabledPtrsPower
}
