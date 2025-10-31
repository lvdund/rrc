package ies

import "rrc/utils"

// MBSFN-AreaInfo-r9 ::= SEQUENCE
// Extensible
type MbsfnAreainfoR9 struct {
	MbsfnAreaidR9           MbsfnAreaidR12
	NonMbsfnregionlength    MbsfnAreainfoR9NonMbsfnregionlength
	NotificationindicatorR9 utils.INTEGER `lb:0,ub:7`
	McchConfigR9            MbsfnAreainfoR9McchConfigR9
}
