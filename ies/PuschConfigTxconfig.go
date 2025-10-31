package ies

import "rrc/utils"

// PUSCH-Config-txConfig ::= ENUMERATED
type PuschConfigTxconfig struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigTxconfigEnumeratedNothing = iota
	PuschConfigTxconfigEnumeratedCodebook
	PuschConfigTxconfigEnumeratedNoncodebook
)
