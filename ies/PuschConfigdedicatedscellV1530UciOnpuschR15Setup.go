package ies

import "rrc/utils"

// PUSCH-ConfigDedicatedScell-v1530-uci-OnPUSCH-r15-setup ::= SEQUENCE
type PuschConfigdedicatedscellV1530UciOnpuschR15Setup struct {
	BetaoffsetaulR15 utils.INTEGER `lb:0,ub:15`
}
