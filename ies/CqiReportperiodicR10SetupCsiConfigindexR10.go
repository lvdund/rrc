package ies

// CQI-ReportPeriodic-r10-setup-csi-ConfigIndex-r10 ::= CHOICE
const (
	CqiReportperiodicR10SetupCsiConfigindexR10ChoiceNothing = iota
	CqiReportperiodicR10SetupCsiConfigindexR10ChoiceRelease
	CqiReportperiodicR10SetupCsiConfigindexR10ChoiceSetup
)

type CqiReportperiodicR10SetupCsiConfigindexR10 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CqiReportperiodicR10SetupCsiConfigindexR10Setup
}
