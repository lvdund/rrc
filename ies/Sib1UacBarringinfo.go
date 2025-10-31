package ies

// SIB1-uac-BarringInfo ::= SEQUENCE
type Sib1UacBarringinfo struct {
	UacBarringforcommon                       *UacBarringpercatlist
	UacBarringperplmnList                     *UacBarringperplmnList
	UacBarringinfosetlist                     UacBarringinfosetlist
	UacAccesscategory1Selectionassistanceinfo *Sib1UacBarringinfoUacAccesscategory1Selectionassistanceinfo
}
