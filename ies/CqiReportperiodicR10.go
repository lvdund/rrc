package ies

// CQI-ReportPeriodic-r10 ::= CHOICE
const (
	CqiReportperiodicR10ChoiceNothing = iota
	CqiReportperiodicR10ChoiceRelease
	CqiReportperiodicR10ChoiceSetup
)

type CqiReportperiodicR10 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CqiReportperiodicR10Setup
}
