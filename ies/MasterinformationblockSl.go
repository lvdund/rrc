package ies

import "rrc/utils"

// MasterInformationBlock-SL ::= SEQUENCE
type MasterinformationblockSl struct {
	SlBandwidthR12          MasterinformationblockSlSlBandwidthR12
	TddConfigslR12          TddConfigslR12
	DirectframenumberR12    utils.BITSTRING `lb:10,ub:10`
	DirectsubframenumberR12 utils.INTEGER   `lb:0,ub:9`
	IncoverageR12           utils.BOOLEAN
	ReservedR12             utils.BITSTRING `lb:19,ub:19`
}
