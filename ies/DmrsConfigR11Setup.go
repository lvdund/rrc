package ies

import "rrc/utils"

// DMRS-Config-r11-setup ::= SEQUENCE
type DmrsConfigR11Setup struct {
	ScramblingidentityR11  utils.INTEGER `lb:0,ub:503`
	Scramblingidentity2R11 utils.INTEGER `lb:0,ub:503`
}
