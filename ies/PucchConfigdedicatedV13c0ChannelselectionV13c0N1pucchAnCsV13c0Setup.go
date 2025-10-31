package ies

import "rrc/utils"

// PUCCH-ConfigDedicated-v13c0-channelSelection-v13c0-n1PUCCH-AN-CS-v13c0-setup ::= SEQUENCE
type PucchConfigdedicatedV13c0ChannelselectionV13c0N1pucchAnCsV13c0Setup struct {
	N1pucchAnCsListp1V13c0 []utils.INTEGER `lb:2,ub:4`
}
