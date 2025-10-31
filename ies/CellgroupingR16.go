package ies

// CellGrouping-r16 ::= SEQUENCE
type CellgroupingR16 struct {
	McgR16  []Freqbandindicatornr `lb:1,ub:maxBands`
	ScgR16  []Freqbandindicatornr `lb:1,ub:maxBands`
	ModeR16 CellgroupingR16ModeR16
}
