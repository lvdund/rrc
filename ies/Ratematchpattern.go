package ies

// RateMatchPattern ::= SEQUENCE
// Extensible
type Ratematchpattern struct {
	Ratematchpatternid    Ratematchpatternid
	Patterntype           *RatematchpatternPatterntype
	Subcarrierspacing     *Subcarrierspacing
	Dummy                 RatematchpatternDummy
	ControlresourcesetR16 *ControlresourcesetidR16 `ext`
}
