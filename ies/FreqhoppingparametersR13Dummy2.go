package ies

import "rrc/utils"

// FreqHoppingParameters-r13-dummy2 ::= CHOICE
const (
	FreqhoppingparametersR13Dummy2ChoiceNothing = iota
	FreqhoppingparametersR13Dummy2ChoiceIntervalFddR13
	FreqhoppingparametersR13Dummy2ChoiceIntervalTddR13
)

type FreqhoppingparametersR13Dummy2 struct {
	Choice         uint64
	IntervalFddR13 *utils.ENUMERATED
	IntervalTddR13 *utils.ENUMERATED
}
