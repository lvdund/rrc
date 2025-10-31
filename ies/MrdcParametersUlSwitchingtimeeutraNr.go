package ies

import "rrc/utils"

// MRDC-Parameters-ul-SwitchingTimeEUTRA-NR ::= ENUMERATED
type MrdcParametersUlSwitchingtimeeutraNr struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersUlSwitchingtimeeutraNrEnumeratedNothing = iota
	MrdcParametersUlSwitchingtimeeutraNrEnumeratedType1
	MrdcParametersUlSwitchingtimeeutraNrEnumeratedType2
)
