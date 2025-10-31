package ies

// CQI-ReportPeriodicSCell-r15 ::= CHOICE
const (
	CqiReportperiodicscellR15ChoiceNothing = iota
	CqiReportperiodicscellR15ChoiceRelease
	CqiReportperiodicscellR15ChoiceSetup
)

type CqiReportperiodicscellR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CqiReportperiodicscellR15Setup
}
