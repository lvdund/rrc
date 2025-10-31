package ies

import "rrc/utils"

// PDSCH-Config-mcs-Table ::= ENUMERATED
type PdschConfigMcsTable struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigMcsTableEnumeratedNothing = iota
	PdschConfigMcsTableEnumeratedQam256
	PdschConfigMcsTableEnumeratedQam64lowse
)
