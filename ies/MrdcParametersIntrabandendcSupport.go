package ies

import "rrc/utils"

// MRDC-Parameters-intraBandENDC-Support ::= ENUMERATED
type MrdcParametersIntrabandendcSupport struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersIntrabandendcSupportEnumeratedNothing = iota
	MrdcParametersIntrabandendcSupportEnumeratedNon_Contiguous
	MrdcParametersIntrabandendcSupportEnumeratedBoth
)
