package ies

// SL-DiscTxRefCarrierDedicated-r13 ::= CHOICE
const (
	SlDisctxrefcarrierdedicatedR13ChoiceNothing = iota
	SlDisctxrefcarrierdedicatedR13ChoicePcell
	SlDisctxrefcarrierdedicatedR13ChoiceScell
)

type SlDisctxrefcarrierdedicatedR13 struct {
	Choice uint64
	Pcell  *struct{}
	Scell  *ScellindexR10
}
