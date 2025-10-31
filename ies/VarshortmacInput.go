package ies

// VarShortMAC-Input ::= SEQUENCE
type VarshortmacInput struct {
	Sourcephyscellid   Physcellid
	Targetcellidentity Cellidentity
	SourceCRnti        RntiValue
}
