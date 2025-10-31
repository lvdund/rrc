package ies

import "rrc/utils"

// CSI-ReportConfig-cqi-Table ::= ENUMERATED
type CsiReportconfigCqiTable struct {
	Value utils.ENUMERATED
}

const (
	CsiReportconfigCqiTableEnumeratedNothing = iota
	CsiReportconfigCqiTableEnumeratedTable1
	CsiReportconfigCqiTableEnumeratedTable2
	CsiReportconfigCqiTableEnumeratedTable3
	CsiReportconfigCqiTableEnumeratedTable4_R17
)
