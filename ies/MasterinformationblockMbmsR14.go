package ies

import "rrc/utils"

// MasterInformationBlock-MBMS-r14 ::= SEQUENCE
type MasterinformationblockMbmsR14 struct {
	DlBandwidthMbmsR14             MasterinformationblockMbmsR14DlBandwidthMbmsR14
	SystemframenumberR14           utils.BITSTRING `lb:6,ub:6`
	AdditionalnonmbsfnsubframesR14 utils.INTEGER   `lb:0,ub:3`
	SemistaticcfiMbmsR16           utils.INTEGER   `lb:0,ub:3`
	Spare                          utils.BITSTRING `lb:11,ub:11`
}
