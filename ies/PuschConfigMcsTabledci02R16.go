package ies

import "rrc/utils"

// PUSCH-Config-mcs-TableDCI-0-2-r16 ::= ENUMERATED
type PuschConfigMcsTabledci02R16 struct {
	Value utils.ENUMERATED
}

const (
	PuschConfigMcsTabledci02R16EnumeratedNothing = iota
	PuschConfigMcsTabledci02R16EnumeratedQam256
	PuschConfigMcsTabledci02R16EnumeratedQam64lowse
)
