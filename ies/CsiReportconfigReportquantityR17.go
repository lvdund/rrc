package ies

// CSI-ReportConfig-reportQuantity-r17 ::= CHOICE
const (
	CsiReportconfigReportquantityR17ChoiceNothing = iota
	CsiReportconfigReportquantityR17ChoiceCriRsrpIndexR17
	CsiReportconfigReportquantityR17ChoiceSsbIndexRsrpIndexR17
	CsiReportconfigReportquantityR17ChoiceCriSinrIndexR17
	CsiReportconfigReportquantityR17ChoiceSsbIndexSinrIndexR17
)

type CsiReportconfigReportquantityR17 struct {
	Choice               uint64
	CriRsrpIndexR17      *struct{}
	SsbIndexRsrpIndexR17 *struct{}
	CriSinrIndexR17      *struct{}
	SsbIndexSinrIndexR17 *struct{}
}
