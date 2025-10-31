package ies

import "rrc/utils"

// EPDCCH-Config-r11-config-r11-setup ::= SEQUENCE
type EpdcchConfigR11ConfigR11Setup struct {
	SubframepatternconfigR11  *EpdcchConfigR11ConfigR11SetupSubframepatternconfigR11
	StartsymbolR11            *utils.INTEGER `lb:0,ub:4`
	SetconfigtoreleaselistR11 *EpdcchSetconfigtoreleaselistR11
	SetconfigtoaddmodlistR11  *EpdcchSetconfigtoaddmodlistR11
}
