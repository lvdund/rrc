package ies

import "rrc/utils"

// ReestabUE-Identity ::= SEQUENCE
type ReestabueIdentity struct {
	CRnti      CRnti
	Physcellid Physcellid
	ShortmacI  ShortmacI
}
