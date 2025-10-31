package ies

import "rrc/utils"

// PCCH-Config-ns ::= ENUMERATED
type PcchConfigNs struct {
	Value utils.ENUMERATED
}

const (
	PcchConfigNsEnumeratedNothing = iota
	PcchConfigNsEnumeratedFour
	PcchConfigNsEnumeratedTwo
	PcchConfigNsEnumeratedOne
)
