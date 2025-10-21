package ies

import "rrc/utils"

// NPDCCH-SC-MCCH-Config-NB-r14 ::= SEQUENCE
type NpdcchScMcchConfigNbR14 struct {
	NpdcchNumrepetitionsScMcchR14 utils.ENUMERATED
	NpdcchStartsfScMcchR14        utils.ENUMERATED
	NpdcchOffsetScMcchR14         utils.ENUMERATED
}
