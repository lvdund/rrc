package ies

import "rrc/utils"

// RACH-ConfigCommon-prach-RootSequenceIndex ::= CHOICE
const (
	RachConfigcommonPrachRootsequenceindexChoiceNothing = iota
	RachConfigcommonPrachRootsequenceindexChoiceL839
	RachConfigcommonPrachRootsequenceindexChoiceL139
)

type RachConfigcommonPrachRootsequenceindex struct {
	Choice uint64
	L839   *utils.INTEGER `lb:0,ub:837`
	L139   *utils.INTEGER `lb:0,ub:137`
}
