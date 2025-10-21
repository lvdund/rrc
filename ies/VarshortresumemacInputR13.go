package ies

import "rrc/utils"

// VarShortResumeMAC-Input-r13 ::= SEQUENCE
type VarshortresumemacInputR13 struct {
	CellidentityR13        Cellidentity
	PhyscellidR13          Physcellid
	CRntiR13               CRnti
	ResumediscriminatorR13 utils.BITSTRING
}
