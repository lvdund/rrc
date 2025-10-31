package ies

import "rrc/utils"

// PDSCH-Config-mcs-Table-r17 ::= ENUMERATED
type PdschConfigMcsTableR17 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigMcsTableR17EnumeratedNothing = iota
	PdschConfigMcsTableR17EnumeratedQam1024
)
