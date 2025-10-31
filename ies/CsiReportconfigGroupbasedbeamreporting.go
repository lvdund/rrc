package ies

// CSI-ReportConfig-groupBasedBeamReporting ::= CHOICE
const (
	CsiReportconfigGroupbasedbeamreportingChoiceNothing = iota
	CsiReportconfigGroupbasedbeamreportingChoiceEnabled
	CsiReportconfigGroupbasedbeamreportingChoiceDisabled
)

type CsiReportconfigGroupbasedbeamreporting struct {
	Choice   uint64
	Enabled  *struct{}
	Disabled *CsiReportconfigGroupbasedbeamreportingDisabled
}
