package ies

// CSI-ReportConfig-reportQuantity ::= CHOICE
const (
	CsiReportconfigReportquantityChoiceNothing = iota
	CsiReportconfigReportquantityChoiceNone
	CsiReportconfigReportquantityChoiceCriRiPmiCqi
	CsiReportconfigReportquantityChoiceCriRiI1
	CsiReportconfigReportquantityChoiceCriRiI1Cqi
	CsiReportconfigReportquantityChoiceCriRiCqi
	CsiReportconfigReportquantityChoiceCriRsrp
	CsiReportconfigReportquantityChoiceSsbIndexRsrp
	CsiReportconfigReportquantityChoiceCriRiLiPmiCqi
)

type CsiReportconfigReportquantity struct {
	Choice        uint64
	None          *struct{}
	CriRiPmiCqi   *struct{}
	CriRiI1       *struct{}
	CriRiI1Cqi    *CsiReportconfigReportquantityCriRiI1Cqi
	CriRiCqi      *struct{}
	CriRsrp       *struct{}
	SsbIndexRsrp  *struct{}
	CriRiLiPmiCqi *struct{}
}
