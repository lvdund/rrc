package ies

import "rrc/utils"

// UL-256QAM-perCC-Info-r14-ul-256QAM-perCC-r14 ::= ENUMERATED
type Ul256qamPerccInfoR14Ul256qamPerccR14 struct {
	Value utils.ENUMERATED
}

const (
	Ul256qamPerccInfoR14Ul256qamPerccR14EnumeratedNothing = iota
	Ul256qamPerccInfoR14Ul256qamPerccR14EnumeratedSupported
)
