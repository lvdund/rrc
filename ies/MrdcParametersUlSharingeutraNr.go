package ies

import "rrc/utils"

// MRDC-Parameters-ul-SharingEUTRA-NR ::= ENUMERATED
type MrdcParametersUlSharingeutraNr struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersUlSharingeutraNrEnumeratedNothing = iota
	MrdcParametersUlSharingeutraNrEnumeratedTdm
	MrdcParametersUlSharingeutraNrEnumeratedFdm
	MrdcParametersUlSharingeutraNrEnumeratedBoth
)
