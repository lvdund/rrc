package ies

import "rrc/utils"

// PCCH-Config-NB-r13-defaultPagingCycle-r13 ::= ENUMERATED
type PcchConfigNbR13DefaultpagingcycleR13 struct {
	Value utils.ENUMERATED
}

const (
	PcchConfigNbR13DefaultpagingcycleR13EnumeratedNothing = iota
	PcchConfigNbR13DefaultpagingcycleR13EnumeratedRf128
	PcchConfigNbR13DefaultpagingcycleR13EnumeratedRf256
	PcchConfigNbR13DefaultpagingcycleR13EnumeratedRf512
	PcchConfigNbR13DefaultpagingcycleR13EnumeratedRf1024
)
