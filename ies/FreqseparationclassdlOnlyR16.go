package ies

import "rrc/utils"

// FreqSeparationClassDL-Only-r16 ::= ENUMERATED
type FreqseparationclassdlOnlyR16 struct {
	Value utils.ENUMERATED
}

const (
	FreqseparationclassdlOnlyR16EnumeratedNothing = iota
	FreqseparationclassdlOnlyR16EnumeratedMhz200
	FreqseparationclassdlOnlyR16EnumeratedMhz400
	FreqseparationclassdlOnlyR16EnumeratedMhz600
	FreqseparationclassdlOnlyR16EnumeratedMhz800
	FreqseparationclassdlOnlyR16EnumeratedMhz1000
	FreqseparationclassdlOnlyR16EnumeratedMhz1200
)
