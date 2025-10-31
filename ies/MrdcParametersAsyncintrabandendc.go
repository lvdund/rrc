package ies

import "rrc/utils"

// MRDC-Parameters-asyncIntraBandENDC ::= ENUMERATED
type MrdcParametersAsyncintrabandendc struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersAsyncintrabandendcEnumeratedNothing = iota
	MrdcParametersAsyncintrabandendcEnumeratedSupported
)
