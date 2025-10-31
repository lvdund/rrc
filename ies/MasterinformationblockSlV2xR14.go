package ies

import "rrc/utils"

// MasterInformationBlock-SL-V2X-r14 ::= SEQUENCE
type MasterinformationblockSlV2xR14 struct {
	SlBandwidthR14          MasterinformationblockSlV2xR14SlBandwidthR14
	TddConfigslR14          TddConfigslR12
	DirectframenumberR14    utils.BITSTRING `lb:10,ub:10`
	DirectsubframenumberR14 utils.INTEGER   `lb:0,ub:9`
	IncoverageR14           utils.BOOLEAN
	ReservedR14             utils.BITSTRING `lb:27,ub:27`
}
