package ies

import "rrc/utils"

// DRX-Config-setup-shortDRX ::= SEQUENCE
type DrxConfigSetupShortdrx struct {
	ShortdrxCycle      DrxConfigSetupShortdrxShortdrxCycle
	Drxshortcycletimer utils.INTEGER `lb:0,ub:16`
}
