package ies

// MeasTriggerQuantity ::= CHOICE
const (
	MeastriggerquantityChoiceNothing = iota
	MeastriggerquantityChoiceRsrp
	MeastriggerquantityChoiceRsrq
	MeastriggerquantityChoiceSinr
)

type Meastriggerquantity struct {
	Choice uint64
	Rsrp   *RsrpRange
	Rsrq   *RsrqRange
	Sinr   *SinrRange
}
