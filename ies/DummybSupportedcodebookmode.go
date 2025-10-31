package ies

import "rrc/utils"

// DummyB-supportedCodebookMode ::= ENUMERATED
type DummybSupportedcodebookmode struct {
	Value utils.ENUMERATED
}

const (
	DummybSupportedcodebookmodeEnumeratedNothing = iota
	DummybSupportedcodebookmodeEnumeratedMode1
	DummybSupportedcodebookmodeEnumeratedMode1andmode2
)
