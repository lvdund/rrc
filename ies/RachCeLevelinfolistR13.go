package ies

import "rrc/utils"

// RACH-CE-LevelInfoList-r13 ::= SEQUENCE OF RACH-CE-LevelInfo-r13
// SIZE (1..maxCE-Level-r13)
type RachCeLevelinfolistR13 struct {
	Value utils.Sequence[RachCeLevelinfoR13]
}
