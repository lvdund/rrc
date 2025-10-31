package ies

import "rrc/utils"

// PDCP-Config-moreThanTwoRLC-DRB-r16 ::= SEQUENCE
type PdcpConfigMorethantworlcDrbR16 struct {
	SplitsecondarypathR16 *Logicalchannelidentity
	DuplicationstateR16   *[]utils.BOOLEAN `lb:3,ub:3`
}
