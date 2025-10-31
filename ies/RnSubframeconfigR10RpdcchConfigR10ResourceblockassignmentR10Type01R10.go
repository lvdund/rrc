package ies

import "rrc/utils"

// RN-SubframeConfig-r10-rpdcch-Config-r10-resourceBlockAssignment-r10-type01-r10 ::= CHOICE
const (
	RnSubframeconfigR10RpdcchConfigR10ResourceblockassignmentR10Type01R10ChoiceNothing = iota
	RnSubframeconfigR10RpdcchConfigR10ResourceblockassignmentR10Type01R10ChoiceNrb6R10
	RnSubframeconfigR10RpdcchConfigR10ResourceblockassignmentR10Type01R10ChoiceNrb15R10
	RnSubframeconfigR10RpdcchConfigR10ResourceblockassignmentR10Type01R10ChoiceNrb25R10
	RnSubframeconfigR10RpdcchConfigR10ResourceblockassignmentR10Type01R10ChoiceNrb50R10
	RnSubframeconfigR10RpdcchConfigR10ResourceblockassignmentR10Type01R10ChoiceNrb75R10
	RnSubframeconfigR10RpdcchConfigR10ResourceblockassignmentR10Type01R10ChoiceNrb100R10
)

type RnSubframeconfigR10RpdcchConfigR10ResourceblockassignmentR10Type01R10 struct {
	Choice    uint64
	Nrb6R10   *utils.BITSTRING `lb:6,ub:6`
	Nrb15R10  *utils.BITSTRING `lb:8,ub:8`
	Nrb25R10  *utils.BITSTRING `lb:13,ub:13`
	Nrb50R10  *utils.BITSTRING `lb:17,ub:17`
	Nrb75R10  *utils.BITSTRING `lb:19,ub:19`
	Nrb100R10 *utils.BITSTRING `lb:25,ub:25`
}
