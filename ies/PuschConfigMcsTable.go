package ies

import "rrc/utils"

// PUSCH-Config-mcs-Table ::= ENUMERATED
type PuschConfigMcsTable struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigMcsTableEnumeratedNothing = iota
	PuschConfigMcsTableEnumeratedQam256
	PuschConfigMcsTableEnumeratedQam64lowse
)
