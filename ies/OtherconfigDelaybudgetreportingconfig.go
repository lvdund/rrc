package ies

import "rrc/utils"

// OtherConfig-delayBudgetReportingConfig ::= CHOICE
const (
	OtherconfigDelaybudgetreportingconfigChoiceNothing = iota
	OtherconfigDelaybudgetreportingconfigChoiceRelease
	OtherconfigDelaybudgetreportingconfigChoiceSetup
)

type OtherconfigDelaybudgetreportingconfig struct {
	Choice  uint64
	Release *struct{}
	Setup   *utils.ENUMERATED
}
