package ies

// CellsTriggeredList-Item ::= CHOICE
const (
	CellstriggeredlistItemChoiceNothing = iota
	CellstriggeredlistItemChoicePhyscellideutra
	CellstriggeredlistItemChoicePhyscellidutra
	CellstriggeredlistItemChoicePhyscellidgeran
	CellstriggeredlistItemChoicePhyscellidcdma2000
	CellstriggeredlistItemChoiceWlanIdentifiersR13
	CellstriggeredlistItemChoicePhyscellidnrR15
)

type CellstriggeredlistItem struct {
	Choice             uint64
	Physcellideutra    *Physcellid
	Physcellidutra     *CellstriggeredlistItemPhyscellidutra
	Physcellidgeran    *CellstriggeredlistItemPhyscellidgeran
	Physcellidcdma2000 *Physcellidcdma2000
	WlanIdentifiersR13 *WlanIdentifiersR12
	PhyscellidnrR15    *CellstriggeredlistItemPhyscellidnrR15
}
