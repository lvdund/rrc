package ies

import "rrc/utils"

// RACH-CE-LevelInfo-r13-rar-HoppingConfig-r13 ::= ENUMERATED
type RachCeLevelinfoR13RarHoppingconfigR13 struct {
	Value utils.ENUMERATED
}

const (
	RachCeLevelinfoR13RarHoppingconfigR13EnumeratedNothing = iota
	RachCeLevelinfoR13RarHoppingconfigR13EnumeratedOn
	RachCeLevelinfoR13RarHoppingconfigR13EnumeratedOff
)
