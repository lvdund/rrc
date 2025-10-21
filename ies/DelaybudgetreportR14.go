package ies

import "rrc/utils"

// DelayBudgetReport-r14 ::= CHOICE
type DelaybudgetreportR14 interface {
	isDelaybudgetreportR14()
}

type DelaybudgetreportR14Type1 struct {
	Value utils.ENUMERATED
}

func (*DelaybudgetreportR14Type1) isDelaybudgetreportR14() {}

type DelaybudgetreportR14Type2 struct {
	Value utils.ENUMERATED
}

func (*DelaybudgetreportR14Type2) isDelaybudgetreportR14() {}
