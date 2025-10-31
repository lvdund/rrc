package ies

// MeasTriggerQuantityEUTRA ::= CHOICE
const (
	MeastriggerquantityeutraChoiceNothing = iota
	MeastriggerquantityeutraChoiceRsrp
	MeastriggerquantityeutraChoiceRsrq
	MeastriggerquantityeutraChoiceSinr
)

type Meastriggerquantityeutra struct {
	Choice uint64
	Rsrp   *RsrpRangeeutra
	Rsrq   *RsrqRangeeutra
	Sinr   *SinrRangeeutra
}
