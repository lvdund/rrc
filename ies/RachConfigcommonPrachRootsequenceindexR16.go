package ies

import "rrc/utils"

// RACH-ConfigCommon-prach-RootSequenceIndex-r16 ::= CHOICE
const (
	RachConfigcommonPrachRootsequenceindexR16ChoiceNothing = iota
	RachConfigcommonPrachRootsequenceindexR16ChoiceL571
	RachConfigcommonPrachRootsequenceindexR16ChoiceL1151
)

type RachConfigcommonPrachRootsequenceindexR16 struct {
	Choice uint64
	L571   *utils.INTEGER `lb:0,ub:569`
	L1151  *utils.INTEGER `lb:0,ub:1149`
}
