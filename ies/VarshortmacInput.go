package ies

// VarShortMAC-Input ::= SEQUENCE
type VarshortmacInput struct {
	Cellidentity Cellidentity
	Physcellid   Physcellid
	CRnti        CRnti
}
