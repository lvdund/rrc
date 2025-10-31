package ies

import "rrc/utils"

// PDSCH-Config-rbg-Size ::= ENUMERATED
type PdschConfigRbgSize struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigRbgSizeEnumeratedNothing = iota
	PdschConfigRbgSizeEnumeratedConfig1
	PdschConfigRbgSizeEnumeratedConfig2
)
