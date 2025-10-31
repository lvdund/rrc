package ies

import "rrc/utils"

// CSI-ProcessId-r11 ::= utils.INTEGER (1..maxCSI-Proc-r11)
type CsiProcessidR11 struct {
	Value utils.INTEGER `lb:0,ub:maxCSIProcR11`
}
