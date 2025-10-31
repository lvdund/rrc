package ies

// RateMatchPattern-patternType ::= CHOICE
// Extensible
const (
	RatematchpatternPatterntypeChoiceNothing = iota
	RatematchpatternPatterntypeChoiceBitmaps
	RatematchpatternPatterntypeChoiceControlresourceset
)

type RatematchpatternPatterntype struct {
	Choice             uint64
	Bitmaps            *RatematchpatternPatterntypeBitmaps
	Controlresourceset *Controlresourcesetid
}
