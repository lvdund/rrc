package ies

// CellsTriggeredList-Item ::= CHOICE
const (
	CellstriggeredlistItemChoiceNothing = iota
	CellstriggeredlistItemChoicePhyscellid
	CellstriggeredlistItemChoicePhyscellideutra
	CellstriggeredlistItemChoicePhyscellidutraFddR16
)

type CellstriggeredlistItem struct {
	Choice               uint64
	Physcellid           *Physcellid
	Physcellideutra      *EutraPhyscellid
	PhyscellidutraFddR16 *PhyscellidutraFddR16
}
