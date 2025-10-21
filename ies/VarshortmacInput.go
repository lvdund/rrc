package ies

import "rrc/utils"

// VarShortMAC-Input ::= SEQUENCE
type VarshortmacInput struct {
	Cellidentity Cellidentity
	Physcellid   Physcellid
	CRnti        CRnti
}
