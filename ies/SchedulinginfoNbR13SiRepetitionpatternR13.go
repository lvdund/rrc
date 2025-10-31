package ies

import "rrc/utils"

// SchedulingInfo-NB-r13-si-RepetitionPattern-r13 ::= ENUMERATED
type SchedulinginfoNbR13SiRepetitionpatternR13 struct {
	Value utils.ENUMERATED
}

const (
	SchedulinginfoNbR13SiRepetitionpatternR13EnumeratedNothing = iota
	SchedulinginfoNbR13SiRepetitionpatternR13EnumeratedEvery2ndrf
	SchedulinginfoNbR13SiRepetitionpatternR13EnumeratedEvery4thrf
	SchedulinginfoNbR13SiRepetitionpatternR13EnumeratedEvery8thrf
	SchedulinginfoNbR13SiRepetitionpatternR13EnumeratedEvery16thrf
)
