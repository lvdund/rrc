package ies

// CQI-ReportAperiodic-v1310-setup-aperiodicCSI-Trigger2-r13 ::= CHOICE
const (
	CqiReportaperiodicV1310SetupAperiodiccsiTrigger2R13ChoiceNothing = iota
	CqiReportaperiodicV1310SetupAperiodiccsiTrigger2R13ChoiceRelease
	CqiReportaperiodicV1310SetupAperiodiccsiTrigger2R13ChoiceSetup
)

type CqiReportaperiodicV1310SetupAperiodiccsiTrigger2R13 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CqiReportaperiodicV1310SetupAperiodiccsiTrigger2R13Setup
}
