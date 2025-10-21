package ies

import "rrc/utils"

// SL-AnchorCarrierFreqList-V2X-r14 ::= SEQUENCE OF ARFCN-ValueEUTRA-r9
// SIZE (1..maxFreqV2X-r14)
type SlAnchorcarrierfreqlistV2xR14 struct {
	Value utils.Sequence[ArfcnValueeutraR9]
}
