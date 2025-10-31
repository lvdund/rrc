package ies

import "rrc/utils"

// RN-SubframeConfig-r10-rpdcch-Config-r10-pucch-Config-r10-tdd-fallbackForFormat3 ::= SEQUENCE
type RnSubframeconfigR10RpdcchConfigR10PucchConfigR10TddFallbackforformat3 struct {
	N1pucchAnP0R10 utils.INTEGER  `lb:0,ub:2047`
	N1pucchAnP1R10 *utils.INTEGER `lb:0,ub:2047`
}
