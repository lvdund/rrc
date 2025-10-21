package ies

import "rrc/utils"

// SRB-ToAddModList-NB-r13 ::= SEQUENCE OF SRB-ToAddMod-NB-r13
// SIZE (1)
type SrbToaddmodlistNbR13 struct {
	Value []SrbToaddmodNbR13
}
