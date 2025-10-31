package ies

import "rrc/utils"

// NR-DL-PRS-ResourceID-r17 ::= utils.INTEGER (0..maxNrofPRS-ResourcesPerSet-1-r17)
type NrDlPrsResourceidR17 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofPRSResourcesperset1R17`
}
