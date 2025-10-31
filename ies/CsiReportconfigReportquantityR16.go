package ies

// CSI-ReportConfig-reportQuantity-r16 ::= CHOICE
const (
	CsiReportconfigReportquantityR16ChoiceNothing = iota
	CsiReportconfigReportquantityR16ChoiceCriSinrR16
	CsiReportconfigReportquantityR16ChoiceSsbIndexSinrR16
)

type CsiReportconfigReportquantityR16 struct {
	Choice          uint64
	CriSinrR16      *struct{}
	SsbIndexSinrR16 *struct{}
}
