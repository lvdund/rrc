package ies

import "rrc/utils"

// PCCH-Config-NB-r13 ::= SEQUENCE
type PcchConfigNbR13 struct {
	DefaultpagingcycleR13        utils.ENUMERATED
	NbR13                        utils.ENUMERATED
	NpdcchNumrepetitionpagingR13 utils.ENUMERATED
}
