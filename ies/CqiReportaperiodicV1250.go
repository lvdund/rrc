package ies

// CQI-ReportAperiodic-v1250 ::= CHOICE
const (
	CqiReportaperiodicV1250ChoiceNothing = iota
	CqiReportaperiodicV1250ChoiceRelease
	CqiReportaperiodicV1250ChoiceSetup
)

type CqiReportaperiodicV1250 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CqiReportaperiodicV1250Setup
}
