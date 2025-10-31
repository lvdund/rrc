package ies

import "rrc/utils"

// MRDC-Parameters-ul-TimingAlignmentEUTRA-NR ::= ENUMERATED
type MrdcParametersUlTimingalignmenteutraNr struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersUlTimingalignmenteutraNrEnumeratedNothing = iota
	MrdcParametersUlTimingalignmenteutraNrEnumeratedRequired
)
