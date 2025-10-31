package ies

import "rrc/utils"

// MAC-Parameters-r12-longDRX-Command-r12 ::= ENUMERATED
type MacParametersR12LongdrxCommandR12 struct {
	Value utils.ENUMERATED
}

const (
	MacParametersR12LongdrxCommandR12EnumeratedNothing = iota
	MacParametersR12LongdrxCommandR12EnumeratedSupported
)
