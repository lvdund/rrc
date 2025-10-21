package ies

import "rrc/utils"

// NPDCCH-ConfigDedicated-NB-r13 ::= SEQUENCE
type NpdcchConfigdedicatedNbR13 struct {
	NpdcchNumrepetitionsR13 utils.ENUMERATED
	NpdcchStartsfUssR13     utils.ENUMERATED
	NpdcchOffsetUssR13      utils.ENUMERATED
}
