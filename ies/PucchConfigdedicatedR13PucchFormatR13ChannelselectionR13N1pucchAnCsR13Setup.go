package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-r13-pucch-Format-r13-channelSelection-r13-n1PUCCH-AN-CS-r13-setup ::= SEQUENCE
type PucchConfigdedicatedR13PucchFormatR13ChannelselectionR13N1pucchAnCsR13Setup struct {
	N1pucchAnCsListR13 []N1pucchAnCsR10 `lb:1,ub:2`
	Dummy1             []utils.INTEGER  `lb:2,ub:4`
}
