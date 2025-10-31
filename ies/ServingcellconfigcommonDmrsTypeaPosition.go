package ies

import "rrc/utils"

// ServingCellConfigCommon-dmrs-TypeA-Position ::= ENUMERATED
type ServingcellconfigcommonDmrsTypeaPosition struct {
	Value utils.ENUMERATED
}

const (
	ServingcellconfigcommonDmrsTypeaPositionEnumeratedNothing = iota
	ServingcellconfigcommonDmrsTypeaPositionEnumeratedPos2
	ServingcellconfigcommonDmrsTypeaPositionEnumeratedPos3
)
