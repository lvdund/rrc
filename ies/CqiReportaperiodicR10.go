package ies

// CQI-ReportAperiodic-r10 ::= CHOICE
const (
	CqiReportaperiodicR10ChoiceNothing = iota
	CqiReportaperiodicR10ChoiceRelease
	CqiReportaperiodicR10ChoiceSetup
)

type CqiReportaperiodicR10 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CqiReportaperiodicR10Setup
}
