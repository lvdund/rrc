package ies

import "rrc/utils"

// FreqHoppingParameters-r13-dummy3 ::= CHOICE
const (
	FreqhoppingparametersR13Dummy3ChoiceNothing = iota
	FreqhoppingparametersR13Dummy3ChoiceIntervalFddR13
	FreqhoppingparametersR13Dummy3ChoiceIntervalTddR13
)

type FreqhoppingparametersR13Dummy3 struct {
	Choice         uint64
	IntervalFddR13 *utils.ENUMERATED
	IntervalTddR13 *utils.ENUMERATED
}
