package ies

// MeasTriggerQuantityCLI-r16 ::= CHOICE
const (
	MeastriggerquantitycliR16ChoiceNothing = iota
	MeastriggerquantitycliR16ChoiceSrsRsrpR16
	MeastriggerquantitycliR16ChoiceCliRssiR16
)

type MeastriggerquantitycliR16 struct {
	Choice     uint64
	SrsRsrpR16 *SrsRsrpRangeR16
	CliRssiR16 *CliRssiRangeR16
}
