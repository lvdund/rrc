package ies

// MeasResultUTRA-physCellId ::= CHOICE
const (
	MeasresultutraPhyscellidChoiceNothing = iota
	MeasresultutraPhyscellidChoiceFdd
	MeasresultutraPhyscellidChoiceTdd
)

type MeasresultutraPhyscellid struct {
	Choice uint64
	Fdd    *PhyscellidutraFdd
	Tdd    *PhyscellidutraTdd
}
