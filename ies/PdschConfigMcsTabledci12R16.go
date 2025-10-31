package ies

import "rrc/utils"

// PDSCH-Config-mcs-TableDCI-1-2-r16 ::= ENUMERATED
type PdschConfigMcsTabledci12R16 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigMcsTabledci12R16EnumeratedNothing = iota
	PdschConfigMcsTabledci12R16EnumeratedQam256
	PdschConfigMcsTabledci12R16EnumeratedQam64lowse
)
