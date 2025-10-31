package ies

// CQI-ReportPeriodicProcExt-r11-csi-ConfigIndex-r11 ::= CHOICE
const (
	CqiReportperiodicprocextR11CsiConfigindexR11ChoiceNothing = iota
	CqiReportperiodicprocextR11CsiConfigindexR11ChoiceRelease
	CqiReportperiodicprocextR11CsiConfigindexR11ChoiceSetup
)

type CqiReportperiodicprocextR11CsiConfigindexR11 struct {
	Choice  uint64
	Release *struct{}
	Setup   *CqiReportperiodicprocextR11CsiConfigindexR11Setup
}
