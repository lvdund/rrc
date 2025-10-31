package ies

import "rrc/utils"

// DummyA-maxNumberCS-IM-PerCC ::= ENUMERATED
type DummyaMaxnumbercsImPercc struct {
	Value utils.ENUMERATED
}

const (
	DummyaMaxnumbercsImPerccEnumeratedNothing = iota
	DummyaMaxnumbercsImPerccEnumeratedN1
	DummyaMaxnumbercsImPerccEnumeratedN2
	DummyaMaxnumbercsImPerccEnumeratedN4
	DummyaMaxnumbercsImPerccEnumeratedN8
	DummyaMaxnumbercsImPerccEnumeratedN16
	DummyaMaxnumbercsImPerccEnumeratedN32
)
