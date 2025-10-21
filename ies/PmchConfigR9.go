package ies

import "rrc/utils"

// PMCH-Config-r9 ::= SEQUENCE
// Extensible
type PmchConfigR9 struct {
	SfAllocendR9          utils.INTEGER
	DatamcsR9             utils.INTEGER
	MchSchedulingperiodR9 utils.ENUMERATED
}
