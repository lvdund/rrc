package ies

import "rrc/utils"

// MeasParameters-v1530-multipleCellsMeasExtension-r15 ::= ENUMERATED
type MeasparametersV1530MultiplecellsmeasextensionR15 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1530MultiplecellsmeasextensionR15EnumeratedNothing = iota
	MeasparametersV1530MultiplecellsmeasextensionR15EnumeratedSupported
)
