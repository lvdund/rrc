package ies

import "rrc/utils"

// PCCH-Config-v1310 ::= SEQUENCE
type PcchConfigV1310 struct {
	PagingNarrowbandsR13         utils.INTEGER `lb:0,ub:maxAvailNarrowBandsR13`
	MpdcchNumrepetitionPagingR13 PcchConfigV1310MpdcchNumrepetitionPagingR13
	NbV1310                      *PcchConfigV1310NbV1310
}
