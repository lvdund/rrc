package ies

// RateMatchPatternGroup-Item ::= CHOICE
const (
	RatematchpatterngroupItemChoiceNothing = iota
	RatematchpatterngroupItemChoiceCelllevel
	RatematchpatterngroupItemChoiceBwplevel
)

type RatematchpatterngroupItem struct {
	Choice    uint64
	Celllevel *Ratematchpatternid
	Bwplevel  *Ratematchpatternid
}
