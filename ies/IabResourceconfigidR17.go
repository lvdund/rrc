package ies

import "rrc/utils"

// IAB-ResourceConfigID-r17 ::= utils.INTEGER (0..maxNrofIABResourceConfig-1-r17)
type IabResourceconfigidR17 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofIABResourceConfig1R17`
}
