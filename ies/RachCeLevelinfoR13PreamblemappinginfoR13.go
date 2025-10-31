package ies

import "rrc/utils"

// RACH-CE-LevelInfo-r13-preambleMappingInfo-r13 ::= SEQUENCE
type RachCeLevelinfoR13PreamblemappinginfoR13 struct {
	FirstpreambleR13 utils.INTEGER `lb:0,ub:63`
	LastpreambleR13  utils.INTEGER `lb:0,ub:63`
}
