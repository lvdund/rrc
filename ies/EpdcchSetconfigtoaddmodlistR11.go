package ies

// EPDCCH-SetConfigToAddModList-r11 ::= SEQUENCE OF EPDCCH-SetConfig-r11
// SIZE (1..maxEPDCCH-Set-r11)
type EpdcchSetconfigtoaddmodlistR11 struct {
	Value []EpdcchSetconfigR11 `lb:1,ub:maxEPDCCHSetR11`
}
