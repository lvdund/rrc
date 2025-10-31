package ies

import "rrc/utils"

// MBSFN-AreaInfo-r16 ::= SEQUENCE
// Extensible
type MbsfnAreainfoR16 struct {
	MbsfnAreaidR16           MbsfnAreaidR12
	NotificationindicatorR16 utils.INTEGER `lb:0,ub:7`
	McchConfigR16            MbsfnAreainfoR16McchConfigR16
	SubcarrierspacingmbmsR16 MbsfnAreainfoR16SubcarrierspacingmbmsR16
	TimeseparationR16        *MbsfnAreainfoR16TimeseparationR16
}
