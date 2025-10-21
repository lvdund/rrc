package ies

import "rrc/utils"

// PeriodicityStartPos-r16 ::= CHOICE
type PeriodicitystartposR16 interface {
	isPeriodicitystartposR16()
}

type PeriodicitystartposR16Periodicity10ms struct {
	Value struct{}
}

func (*PeriodicitystartposR16Periodicity10ms) isPeriodicitystartposR16() {}

type PeriodicitystartposR16Periodicity20ms struct {
	Value utils.INTEGER
}

func (*PeriodicitystartposR16Periodicity20ms) isPeriodicitystartposR16() {}

type PeriodicitystartposR16Periodicity40ms struct {
	Value utils.INTEGER
}

func (*PeriodicitystartposR16Periodicity40ms) isPeriodicitystartposR16() {}

type PeriodicitystartposR16Periodicity80ms struct {
	Value utils.INTEGER
}

func (*PeriodicitystartposR16Periodicity80ms) isPeriodicitystartposR16() {}

type PeriodicitystartposR16Periodicity160ms struct {
	Value utils.INTEGER
}

func (*PeriodicitystartposR16Periodicity160ms) isPeriodicitystartposR16() {}

type PeriodicitystartposR16Spare3 struct {
	Value struct{}
}

func (*PeriodicitystartposR16Spare3) isPeriodicitystartposR16() {}

type PeriodicitystartposR16Spare2 struct {
	Value struct{}
}

func (*PeriodicitystartposR16Spare2) isPeriodicitystartposR16() {}

type PeriodicitystartposR16Spare1 struct {
	Value struct{}
}

func (*PeriodicitystartposR16Spare1) isPeriodicitystartposR16() {}
