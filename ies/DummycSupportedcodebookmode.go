package ies

import "rrc/utils"

// DummyC-supportedCodebookMode ::= ENUMERATED
type DummycSupportedcodebookmode struct {
	Value utils.ENUMERATED
}

const (
	DummycSupportedcodebookmodeEnumeratedNothing = iota
	DummycSupportedcodebookmodeEnumeratedMode1
	DummycSupportedcodebookmodeEnumeratedMode2
	DummycSupportedcodebookmodeEnumeratedBoth
)
