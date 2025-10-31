package ies

import "rrc/utils"

// PMCH-Config-r9 ::= SEQUENCE
// Extensible
type PmchConfigR9 struct {
	SfAllocendR9          utils.INTEGER `lb:0,ub:1535`
	DatamcsR9             utils.INTEGER `lb:0,ub:28`
	MchSchedulingperiodR9 PmchConfigR9MchSchedulingperiodR9
}
