package ies

import "rrc/utils"

// PCCH-Config-NB-r14 ::= SEQUENCE
// Extensible
type PcchConfigNbR14 struct {
	NpdcchNumrepetitionpagingR14 *utils.ENUMERATED
	PagingweightR14              PagingweightNbR14
}
