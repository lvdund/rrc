package ies

import "rrc/utils"

// PMCH-Config-r12 ::= SEQUENCE
// Extensible
type PmchConfigR12 struct {
	SfAllocendR12          utils.INTEGER
	DatamcsR12             interface{}
	MchSchedulingperiodR12 utils.ENUMERATED
}
