package ies

import "rrc/utils"

// SL-Parameters-v1310-discInterFreqTx-r13 ::= ENUMERATED
type SlParametersV1310DiscinterfreqtxR13 struct {
	Value utils.ENUMERATED
}

const (
	SlParametersV1310DiscinterfreqtxR13EnumeratedNothing = iota
	SlParametersV1310DiscinterfreqtxR13EnumeratedSupported
)
