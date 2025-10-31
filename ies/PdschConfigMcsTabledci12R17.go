package ies

import "rrc/utils"

// PDSCH-Config-mcs-TableDCI-1-2-r17 ::= ENUMERATED
type PdschConfigMcsTabledci12R17 struct {
	Value utils.ENUMERATED
}

const (
	PdschConfigMcsTabledci12R17EnumeratedNothing = iota
	PdschConfigMcsTabledci12R17EnumeratedQam1024
)
