package ies

// CQI-ReportAperiodic-v1310 ::= CHOICE
const (
	CqiReportaperiodicV1310ChoiceNothing = iota
	CqiReportaperiodicV1310ChoiceRelease
	CqiReportaperiodicV1310ChoiceSetup
)

type CqiReportaperiodicV1310 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CqiReportaperiodicV1310Setup
}
