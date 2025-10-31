package ies

import "rrc/utils"

// FreqHoppingParameters-r13-interval-ULHoppingConfigCommonModeB-r13 ::= CHOICE
const (
	FreqhoppingparametersR13IntervalUlhoppingconfigcommonmodebR13ChoiceNothing = iota
	FreqhoppingparametersR13IntervalUlhoppingconfigcommonmodebR13ChoiceIntervalFddR13
	FreqhoppingparametersR13IntervalUlhoppingconfigcommonmodebR13ChoiceIntervalTddR13
)

type FreqhoppingparametersR13IntervalUlhoppingconfigcommonmodebR13 struct {
	Choice         uint64
	IntervalFddR13 *utils.ENUMERATED
	IntervalTddR13 *utils.ENUMERATED
}
