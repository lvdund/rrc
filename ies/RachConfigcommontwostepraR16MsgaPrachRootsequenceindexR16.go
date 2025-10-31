package ies

import "rrc/utils"

// RACH-ConfigCommonTwoStepRA-r16-msgA-PRACH-RootSequenceIndex-r16 ::= CHOICE
const (
	RachConfigcommontwostepraR16MsgaPrachRootsequenceindexR16ChoiceNothing = iota
	RachConfigcommontwostepraR16MsgaPrachRootsequenceindexR16ChoiceL839
	RachConfigcommontwostepraR16MsgaPrachRootsequenceindexR16ChoiceL139
	RachConfigcommontwostepraR16MsgaPrachRootsequenceindexR16ChoiceL571
	RachConfigcommontwostepraR16MsgaPrachRootsequenceindexR16ChoiceL1151
)

type RachConfigcommontwostepraR16MsgaPrachRootsequenceindexR16 struct {
	Choice uint64
	L839   *utils.INTEGER `lb:0,ub:837`
	L139   *utils.INTEGER `lb:0,ub:137`
	L571   *utils.INTEGER `lb:0,ub:569`
	L1151  *utils.INTEGER `lb:0,ub:1149`
}
