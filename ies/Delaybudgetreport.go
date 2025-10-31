package ies

import "rrc/utils"

// DelayBudgetReport ::= CHOICE
// Extensible
const (
	DelaybudgetreportChoiceNothing = iota
	DelaybudgetreportChoiceType1
)

type Delaybudgetreport struct {
	Choice uint64
	Type1  *utils.ENUMERATED
}
