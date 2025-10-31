package ies

import "rrc/utils"

// RN-SubframeConfig-r10-rpdcch-Config-r10-resourceBlockAssignment-r10-type2-r10 ::= CHOICE
const (
	RnSubframeconfigR10RpdcchConfigR10ResourceblockassignmentR10Type2R10ChoiceNothing = iota
	RnSubframeconfigR10RpdcchConfigR10ResourceblockassignmentR10Type2R10ChoiceNrb6R10
	RnSubframeconfigR10RpdcchConfigR10ResourceblockassignmentR10Type2R10ChoiceNrb15R10
	RnSubframeconfigR10RpdcchConfigR10ResourceblockassignmentR10Type2R10ChoiceNrb25R10
	RnSubframeconfigR10RpdcchConfigR10ResourceblockassignmentR10Type2R10ChoiceNrb50R10
	RnSubframeconfigR10RpdcchConfigR10ResourceblockassignmentR10Type2R10ChoiceNrb75R10
	RnSubframeconfigR10RpdcchConfigR10ResourceblockassignmentR10Type2R10ChoiceNrb100R10
)

type RnSubframeconfigR10RpdcchConfigR10ResourceblockassignmentR10Type2R10 struct {
	Choice    uint64
	Nrb6R10   *utils.BITSTRING `lb:5,ub:5`
	Nrb15R10  *utils.BITSTRING `lb:7,ub:7`
	Nrb25R10  *utils.BITSTRING `lb:9,ub:9`
	Nrb50R10  *utils.BITSTRING `lb:11,ub:11`
	Nrb75R10  *utils.BITSTRING `lb:12,ub:12`
	Nrb100R10 *utils.BITSTRING `lb:13,ub:13`
}
