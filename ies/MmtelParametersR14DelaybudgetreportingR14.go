package ies

import "rrc/utils"

// MMTEL-Parameters-r14-delayBudgetReporting-r14 ::= ENUMERATED
type MmtelParametersR14DelaybudgetreportingR14 struct {
	Value utils.ENUMERATED
}

const (
	MmtelParametersR14DelaybudgetreportingR14EnumeratedNothing = iota
	MmtelParametersR14DelaybudgetreportingR14EnumeratedSupported
)
