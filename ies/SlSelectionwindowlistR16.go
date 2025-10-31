package ies

// SL-SelectionWindowList-r16 ::= SEQUENCE OF SL-SelectionWindowConfig-r16
// SIZE (8)
type SlSelectionwindowlistR16 struct {
	Value []SlSelectionwindowconfigR16 `lb:8,ub:8`
}
