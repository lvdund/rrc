package ies

// EPDCCH-SetConfigToReleaseList-r11 ::= SEQUENCE OF EPDCCH-SetConfigId-r11
// SIZE (1..maxEPDCCH-Set-r11)
type EpdcchSetconfigtoreleaselistR11 struct {
	Value []EpdcchSetconfigidR11 `lb:1,ub:maxEPDCCHSetR11`
}
