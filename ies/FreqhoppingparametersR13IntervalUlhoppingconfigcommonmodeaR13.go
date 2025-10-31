package ies

import "rrc/utils"

// FreqHoppingParameters-r13-interval-ULHoppingConfigCommonModeA-r13 ::= CHOICE
const (
	FreqhoppingparametersR13IntervalUlhoppingconfigcommonmodeaR13ChoiceNothing = iota
	FreqhoppingparametersR13IntervalUlhoppingconfigcommonmodeaR13ChoiceIntervalFddR13
	FreqhoppingparametersR13IntervalUlhoppingconfigcommonmodeaR13ChoiceIntervalTddR13
)

type FreqhoppingparametersR13IntervalUlhoppingconfigcommonmodeaR13 struct {
	Choice         uint64
	IntervalFddR13 *utils.ENUMERATED
	IntervalTddR13 *utils.ENUMERATED
}
