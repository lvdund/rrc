package ies

// VarResumeMAC-Input ::= SEQUENCE
type VarresumemacInput struct {
	Sourcephyscellid   Physcellid
	Targetcellidentity Cellidentity
	SourceCRnti        RntiValue
}
