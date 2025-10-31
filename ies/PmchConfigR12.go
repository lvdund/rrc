package ies

import "rrc/utils"

// PMCH-Config-r12 ::= SEQUENCE
// Extensible
type PmchConfigR12 struct {
	SfAllocendR12          utils.INTEGER `lb:0,ub:1535`
	DatamcsR12             PmchConfigR12DatamcsR12
	MchSchedulingperiodR12 PmchConfigR12MchSchedulingperiodR12
}
