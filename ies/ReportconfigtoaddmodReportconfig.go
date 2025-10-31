package ies

// ReportConfigToAddMod-reportConfig ::= CHOICE
// Extensible
const (
	ReportconfigtoaddmodReportconfigChoiceNothing = iota
	ReportconfigtoaddmodReportconfigChoiceReportconfignr
	ReportconfigtoaddmodReportconfigChoiceReportconfiginterrat
	ReportconfigtoaddmodReportconfigChoiceReportconfignrSlR16
)

type ReportconfigtoaddmodReportconfig struct {
	Choice               uint64
	Reportconfignr       *Reportconfignr
	Reportconfiginterrat *Reportconfiginterrat
	ReportconfignrSlR16  *ReportconfignrSlR16
}
