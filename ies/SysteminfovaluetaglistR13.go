package ies

import "rrc/utils"

// SystemInfoValueTagList-r13 ::= SEQUENCE OF SystemInfoValueTagSI-r13
// SIZE (1..maxSI-Message)
type SysteminfovaluetaglistR13 struct {
	Value utils.Sequence[SysteminfovaluetagsiR13]
}
