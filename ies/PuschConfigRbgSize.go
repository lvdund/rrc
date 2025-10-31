package ies

import "rrc/utils"

// PUSCH-Config-rbg-Size ::= ENUMERATED
type PuschConfigRbgSize struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigRbgSizeEnumeratedNothing = iota
	PuschConfigRbgSizeEnumeratedConfig2
)
