package ies

import "rrc/utils"

// PCCH-Config-v1310 ::= SEQUENCE
type PcchConfigV1310 struct {
	PagingNarrowbandsR13         utils.INTEGER
	MpdcchNumrepetitionPagingR13 utils.ENUMERATED
	NbV1310                      *utils.ENUMERATED
}
