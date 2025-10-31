package ies

// CellsTriggeredList-Item-physCellIdUTRA ::= CHOICE
const (
	CellstriggeredlistItemPhyscellidutraChoiceNothing = iota
	CellstriggeredlistItemPhyscellidutraChoiceFdd
	CellstriggeredlistItemPhyscellidutraChoiceTdd
)

type CellstriggeredlistItemPhyscellidutra struct {
	Choice uint64
	Fdd    *PhyscellidutraFdd
	Tdd    *PhyscellidutraTdd
}
