package ies

import "rrc/utils"

// DRX-Info-shortDRX ::= SEQUENCE
type DrxInfoShortdrx struct {
	DrxShortcycle      DrxInfoShortdrxDrxShortcycle
	DrxShortcycletimer utils.INTEGER `lb:0,ub:16`
}
