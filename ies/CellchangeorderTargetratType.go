package ies

// CellChangeOrder-targetRAT-Type ::= CHOICE
// Extensible
const (
	CellchangeorderTargetratTypeChoiceNothing = iota
	CellchangeorderTargetratTypeChoiceGeran
)

type CellchangeorderTargetratType struct {
	Choice uint64
	Geran  *CellchangeorderTargetratTypeGeran
}
