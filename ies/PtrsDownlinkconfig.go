package ies

import "rrc/utils"

// PTRS-DownlinkConfig ::= SEQUENCE
// Extensible
type PtrsDownlinkconfig struct {
	Frequencydensity      *[]utils.INTEGER `lb:2,ub:2`
	Timedensity           *[]utils.INTEGER `lb:3,ub:3`
	EpreRatio             *utils.INTEGER   `lb:0,ub:3`
	Resourceelementoffset *PtrsDownlinkconfigResourceelementoffset
	MaxnrofportsR16       *PtrsDownlinkconfigMaxnrofportsR16 `ext`
}
