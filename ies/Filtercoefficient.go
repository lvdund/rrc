package ies

import "rrc/utils"

// FilterCoefficient ::= utils.ENUMERATED // Extensible
type Filtercoefficient struct {
	Value utils.ENUMERATED
}

const (
	FiltercoefficientEnumeratedNothing = iota
	FiltercoefficientEnumeratedFc0
	FiltercoefficientEnumeratedFc1
	FiltercoefficientEnumeratedFc2
	FiltercoefficientEnumeratedFc3
	FiltercoefficientEnumeratedFc4
	FiltercoefficientEnumeratedFc5
	FiltercoefficientEnumeratedFc6
	FiltercoefficientEnumeratedFc7
	FiltercoefficientEnumeratedFc8
	FiltercoefficientEnumeratedFc9
	FiltercoefficientEnumeratedFc11
	FiltercoefficientEnumeratedFc13
	FiltercoefficientEnumeratedFc15
	FiltercoefficientEnumeratedFc17
	FiltercoefficientEnumeratedFc19
	FiltercoefficientEnumeratedSpare1
)
