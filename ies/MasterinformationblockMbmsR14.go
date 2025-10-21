package ies

import "rrc/utils"

// MasterInformationBlock-MBMS-r14 ::= SEQUENCE
type MasterinformationblockMbmsR14 struct {
	DlBandwidthMbmsR14             utils.ENUMERATED
	SystemframenumberR14           utils.BITSTRING
	AdditionalnonmbsfnsubframesR14 utils.INTEGER
	SemistaticcfiMbmsR16           utils.INTEGER
	Spare                          utils.BITSTRING
}
