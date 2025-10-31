package ies

import "rrc/utils"

// RN-SubframeConfig-r10-rpdcch-Config-r10 ::= SEQUENCE
// Extensible
type RnSubframeconfigR10RpdcchConfigR10 struct {
	ResourceallocationtypeR10  RnSubframeconfigR10RpdcchConfigR10ResourceallocationtypeR10
	ResourceblockassignmentR10 RnSubframeconfigR10RpdcchConfigR10ResourceblockassignmentR10
	DemodulationrsR10          RnSubframeconfigR10RpdcchConfigR10DemodulationrsR10
	PdschStartR10              utils.INTEGER `lb:0,ub:3`
	PucchConfigR10             *RnSubframeconfigR10RpdcchConfigR10PucchConfigR10
}
