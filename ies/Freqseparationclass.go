package ies

import "rrc/utils"

// FreqSeparationClass ::= utils.ENUMERATED // Extensible
type Freqseparationclass struct {
	Value utils.ENUMERATED
}

const (
	FreqseparationclassEnumeratedNothing = iota
	FreqseparationclassEnumeratedMhz800
	FreqseparationclassEnumeratedMhz1200
	FreqseparationclassEnumeratedMhz1400
	FreqseparationclassEnumeratedMhz400_V1650
	FreqseparationclassEnumeratedMhz600_V1650
)
