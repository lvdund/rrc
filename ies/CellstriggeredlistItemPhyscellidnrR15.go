package ies

// CellsTriggeredList-Item-physCellIdNR-r15 ::= SEQUENCE
type CellstriggeredlistItemPhyscellidnrR15 struct {
	Carrierfreq    ArfcnValuenrR15
	Physcellid     PhyscellidnrR15
	RsIndexlistR15 *SsbIndexlistR15
}
