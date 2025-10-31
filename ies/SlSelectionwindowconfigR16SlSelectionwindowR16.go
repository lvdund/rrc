package ies

import "rrc/utils"

// SL-SelectionWindowConfig-r16-sl-SelectionWindow-r16 ::= ENUMERATED
type SlSelectionwindowconfigR16SlSelectionwindowR16 struct {
	Value utils.ENUMERATED
}

const (
	SlSelectionwindowconfigR16SlSelectionwindowR16EnumeratedNothing = iota
	SlSelectionwindowconfigR16SlSelectionwindowR16EnumeratedN1
	SlSelectionwindowconfigR16SlSelectionwindowR16EnumeratedN5
	SlSelectionwindowconfigR16SlSelectionwindowR16EnumeratedN10
	SlSelectionwindowconfigR16SlSelectionwindowR16EnumeratedN20
)
