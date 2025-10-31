package ies

// CQI-ReportPeriodic ::= CHOICE
const (
	CqiReportperiodicChoiceNothing = iota
	CqiReportperiodicChoiceRelease
	CqiReportperiodicChoiceSetup
)

type CqiReportperiodic struct {
	Choice  uint64
	Release *struct{}
	Setup   *CqiReportperiodicSetup
}
