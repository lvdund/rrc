package ies

import "rrc/utils"

// DummyC-supportedNumberPanels ::= ENUMERATED
type DummycSupportednumberpanels struct {
	Value utils.ENUMERATED
}

const (
	DummycSupportednumberpanelsEnumeratedNothing = iota
	DummycSupportednumberpanelsEnumeratedN2
	DummycSupportednumberpanelsEnumeratedN4
)
