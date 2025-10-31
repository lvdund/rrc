package ies

import "rrc/utils"

// FreqSeparationClassDL-v1620 ::= ENUMERATED
type FreqseparationclassdlV1620 struct {
	Value utils.ENUMERATED
}

const (
	FreqseparationclassdlV1620EnumeratedNothing = iota
	FreqseparationclassdlV1620EnumeratedMhz1000
	FreqseparationclassdlV1620EnumeratedMhz1600
	FreqseparationclassdlV1620EnumeratedMhz1800
	FreqseparationclassdlV1620EnumeratedMhz2000
	FreqseparationclassdlV1620EnumeratedMhz2200
	FreqseparationclassdlV1620EnumeratedMhz2400
)
