package ies

// CSI-ReportConfig-reportConfigType ::= CHOICE
const (
	CsiReportconfigReportconfigtypeChoiceNothing = iota
	CsiReportconfigReportconfigtypeChoicePeriodic
	CsiReportconfigReportconfigtypeChoiceSemipersistentonpucch
	CsiReportconfigReportconfigtypeChoiceSemipersistentonpusch
	CsiReportconfigReportconfigtypeChoiceAperiodic
)

type CsiReportconfigReportconfigtype struct {
	Choice                uint64
	Periodic              *CsiReportconfigReportconfigtypePeriodic
	Semipersistentonpucch *CsiReportconfigReportconfigtypeSemipersistentonpucch
	Semipersistentonpusch *CsiReportconfigReportconfigtypeSemipersistentonpusch
	Aperiodic             *CsiReportconfigReportconfigtypeAperiodic
}
