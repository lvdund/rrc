package ies

// TwoPUCCH-Grp-ConfigParams-r16 ::= SEQUENCE
type TwopucchGrpConfigparamsR16 struct {
	PucchGroupmappingR16 PucchGrpCarriertypesR16
	PucchTxR16           PucchGrpCarriertypesR16
}
