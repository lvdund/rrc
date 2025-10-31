package ies

import "rrc/utils"

// GWUS-NumGroups-NB-r16 ::= ENUMERATED
type GwusNumgroupsNbR16 struct {
	Value utils.ENUMERATED
}

const (
	GwusNumgroupsNbR16EnumeratedNothing = iota
	GwusNumgroupsNbR16EnumeratedN1
	GwusNumgroupsNbR16EnumeratedN2
	GwusNumgroupsNbR16EnumeratedN4
	GwusNumgroupsNbR16EnumeratedN8
)
