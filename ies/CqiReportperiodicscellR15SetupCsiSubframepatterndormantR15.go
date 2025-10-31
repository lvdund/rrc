package ies

// CQI-ReportPeriodicSCell-r15-setup-csi-SubframePatternDormant-r15 ::= CHOICE
const (
	CqiReportperiodicscellR15SetupCsiSubframepatterndormantR15ChoiceNothing = iota
	CqiReportperiodicscellR15SetupCsiSubframepatterndormantR15ChoiceRelease
	CqiReportperiodicscellR15SetupCsiSubframepatterndormantR15ChoiceSetup
)

type CqiReportperiodicscellR15SetupCsiSubframepatterndormantR15 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CqiReportperiodicscellR15SetupCsiSubframepatterndormantR15Setup
}
