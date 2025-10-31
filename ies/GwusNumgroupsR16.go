package ies

import "rrc/utils"

// GWUS-NumGroups-r16 ::= ENUMERATED
type GwusNumgroupsR16 struct {
	Value utils.ENUMERATED
}

const (
	GwusNumgroupsR16EnumeratedNothing = iota
	GwusNumgroupsR16EnumeratedN1
	GwusNumgroupsR16EnumeratedN2
	GwusNumgroupsR16EnumeratedN4
	GwusNumgroupsR16EnumeratedN8
)
