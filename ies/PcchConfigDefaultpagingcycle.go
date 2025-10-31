package ies

import "rrc/utils"

// PCCH-Config-defaultPagingCycle ::= ENUMERATED
type PcchConfigDefaultpagingcycle struct {
	Value utils.ENUMERATED
}

const (
	PcchConfigDefaultpagingcycleEnumeratedNothing = iota
	PcchConfigDefaultpagingcycleEnumeratedRf32
	PcchConfigDefaultpagingcycleEnumeratedRf64
	PcchConfigDefaultpagingcycleEnumeratedRf128
	PcchConfigDefaultpagingcycleEnumeratedRf256
)
