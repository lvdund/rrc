package ies

import "rrc/utils"

// ThresholdUTRA ::= CHOICE
const (
	ThresholdutraChoiceNothing = iota
	ThresholdutraChoiceUtraRscp
	ThresholdutraChoiceUtraEcn0
)

type Thresholdutra struct {
	Choice   uint64
	UtraRscp *utils.INTEGER `lb:-5,ub:91`
	UtraEcn0 *utils.INTEGER `lb:0,ub:49`
}
