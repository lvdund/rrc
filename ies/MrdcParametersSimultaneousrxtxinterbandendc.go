package ies

import "rrc/utils"

// MRDC-Parameters-simultaneousRxTxInterBandENDC ::= ENUMERATED
type MrdcParametersSimultaneousrxtxinterbandendc struct {
	Value utils.ENUMERATED
}

const (
	MrdcParametersSimultaneousrxtxinterbandendcEnumeratedNothing = iota
	MrdcParametersSimultaneousrxtxinterbandendcEnumeratedSupported
)
