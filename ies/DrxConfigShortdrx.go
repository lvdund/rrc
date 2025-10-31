package ies

import "rrc/utils"

// DRX-Config-shortDRX ::= SEQUENCE
type DrxConfigShortdrx struct {
	DrxShortcycle      DrxConfigShortdrxDrxShortcycle
	DrxShortcycletimer utils.INTEGER `lb:0,ub:16`
}
