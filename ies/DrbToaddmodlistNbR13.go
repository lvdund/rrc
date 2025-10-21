package ies

import "rrc/utils"

// DRB-ToAddModList-NB-r13 ::= SEQUENCE OF DRB-ToAddMod-NB-r13
// SIZE (1..maxDRB-NB-r13)
type DrbToaddmodlistNbR13 struct {
	Value []DrbToaddmodNbR13
}
