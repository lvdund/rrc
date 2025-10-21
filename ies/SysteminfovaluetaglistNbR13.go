package ies

import "rrc/utils"

// SystemInfoValueTagList-NB-r13 ::= SEQUENCE OF SystemInfoValueTagSI-r13
// SIZE (1.. maxSI-Message-NB-r13)
type SysteminfovaluetaglistNbR13 struct {
	Value []SysteminfovaluetagsiR13
}
