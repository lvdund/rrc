package ies

import "rrc/utils"

// MeasConfigAppLayerId-r17 ::= utils.INTEGER (0..maxNrofAppLayerMeas-1-r17)
type MeasconfigapplayeridR17 struct {
	Value utils.INTEGER `lb:0,ub:maxNrofAppLayerMeas1R17`
}
