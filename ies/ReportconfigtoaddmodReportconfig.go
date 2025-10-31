package ies

// ReportConfigToAddMod-reportConfig ::= CHOICE
const (
	ReportconfigtoaddmodReportconfigChoiceNothing = iota
	ReportconfigtoaddmodReportconfigChoiceReportconfigeutra
	ReportconfigtoaddmodReportconfigChoiceReportconfiginterrat
)

type ReportconfigtoaddmodReportconfig struct {
	Choice               uint64
	Reportconfigeutra    *Reportconfigeutra
	Reportconfiginterrat *Reportconfiginterrat
}
