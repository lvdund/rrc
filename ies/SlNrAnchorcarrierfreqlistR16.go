package ies

import "rrc/utils"

// SL-NR-AnchorCarrierFreqList-r16 ::= SEQUENCE OF ARFCN-ValueNR-r15
// SIZE (1..maxFreqSL-NR-r16)
type SlNrAnchorcarrierfreqlistR16 struct {
	Value utils.Sequence[ArfcnValuenrR15]
}
