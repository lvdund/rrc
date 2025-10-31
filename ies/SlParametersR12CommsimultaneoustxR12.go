package ies

import "rrc/utils"

// SL-Parameters-r12-commSimultaneousTx-r12 ::= ENUMERATED
type SlParametersR12CommsimultaneoustxR12 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersR12CommsimultaneoustxR12EnumeratedNothing = iota
	SlParametersR12CommsimultaneoustxR12EnumeratedSupported
)
