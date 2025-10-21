package ies

import "rrc/utils"

// MasterInformationBlock-SL-V2X-r14 ::= SEQUENCE
type MasterinformationblockSlV2xR14 struct {
	SlBandwidthR14          utils.ENUMERATED
	TddConfigslR14          TddConfigslR12
	DirectframenumberR14    utils.BITSTRING
	DirectsubframenumberR14 utils.INTEGER
	IncoverageR14           bool
	ReservedR14             utils.BITSTRING
}
